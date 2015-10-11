package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
)

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}

func RoundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}

type Dividable interface {
	IsRestlessDividable(uint32) bool
	GetMn() uint32
}

type Division struct {
	divisor uint32
	mn      uint32
}

func NewDivision(divisor uint32) Dividable {
	division := &Division{divisor, math.MaxUint32/divisor + 1}
	return division
}

func (division *Division) IsRestlessDividable(dividend uint32) bool {
	if dividend == 0 {
		return true
	}
	if dividend < division.divisor {
		return false
	}
	return (dividend * division.mn) < division.mn
}

func (division *Division) GetMn() uint32 {
	return division.mn
}

type DivisionPow2 struct {
	divisor uint32
	mn      uint32
}

func NewDivisionPow2(divisor uint32) Dividable {
	division := &DivisionPow2{divisor, divisor - 1}
	return division
}

func (division *DivisionPow2) IsRestlessDividable(dividend uint32) bool {
	if dividend == 0 {
		return true
	}
	if dividend < division.divisor {
		return false
	}
	return (dividend & division.mn) == 0
}

func (division *DivisionPow2) GetMn() uint32 {
	return division.mn
}

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
			//if RoundPlus(1.0/float64(exactCount)*float64(fastCount), 4) != 1.0 {
			fmt.Println("Divisor:", i, "Fast wrong:", fastWrong, "Zeroremainder counts:", fastCount, "Modulo operation counts:", exactCount, "Difference (%):", 100.0-RoundPlus(100.0/float64(exactCount)*float64(fastCount), 4))
			//}
		}(i)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
}
