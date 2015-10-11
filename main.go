package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup

	var minDividend uint32 = 0
	var maxDividend uint32 = 0xffff // fully tested for 0xffffff, may take a long time

	var minDivisor uint32 = 2    // min is 2 because negative number and zero don't make sense, and 1 means every packet (no selection)
	var maxDivisor uint32 = 1000 // tested until 100000

	var i uint32
	for i = minDivisor; i <= maxDivisor; i++ {
		go func(i uint32) {
			wg.Add(1)
			defer wg.Done()

			var j uint32
			var d Dividable

			if i&(i-1) == 0 {
				d = NewDivisionPow2(i)
			} else {
				d = NewDivision(i)
			}
			var fastCount, exactCount, fastWrong int
			for j = minDividend; j < maxDividend; j++ {
				fast := d.IsRestlessDividable(j)
				exact := (j%i == 0)
				if fast != exact {
					fastWrong++
					if false {
						fmt.Println("Error: ", j, i, d.IsRestlessDividable(j), (j%i == 0), d.GetMn())
						os.Exit(0)
					}
				}
				fastCount += Btoi(fast)
				exactCount += Btoi(exact)
			}
			//if Round(1.0/float64(exactCount)*float64(fastCount), 4) != 1.0 {
			fmt.Println("Divisor:", i, "Fast wrong:", fastWrong, "Zeroremainder counts:", fastCount, "Modulo operation counts:", exactCount, "Difference (%):", 100.0-Round(100.0/float64(exactCount)*float64(fastCount), 4))
			//}
		}(i)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
}
