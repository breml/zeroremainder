package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	flagDividendStart        = flag.Int("dividendstart", 0, "Start of range of dividends (uint32)")
	flagDividendEnd          = flag.Int("dividendend", 0xffff, "End of range of dividends (uint32)")
	flagDivisorStart         = flag.Int("divisorstart", 2, "Start of range of divisors (uint32)")
	flagDivisorEnd           = flag.Int("divisorend", 1000, "End of range of divisors (uint32)")
	flagMaxProcs             = flag.Int("maxprocs", 0, "Value for GOMAXPROCS (value may be reduced, respecting default GOMAXPROCS and number of CPUs)")
	flagOutputAllDivisors    = flag.Bool("outputalldivisors", false, "Output result for all divisors")
	flagOutputAllDifferences = flag.Bool("outputalldifferences", false, "Output every dividend / divisor combination which provides a flase result")
	flagOutputAll            = flag.Bool("outputall", false, "Output every calculation, includes -outputalldivisors and -outputalldifferences; WARNING: very verbose, use only with small ranges for dividend and divisor, maxprocs should be set to 1")
	flagIgnorePow2           = flag.Bool("ignorepow2", true, "for power of 2 divisors the solution is trivial, therefore these divisors may be ignored")
	flagUsePow2              = flag.Bool("usepow2", true, "use bitwise and for power of 2 divisors")
	flagHelp                 = flag.Bool("help", false, "Print help")
)

func main() {
	flag.Parse()

	if *flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	runtime.GOMAXPROCS(MaxParallelism(*flagMaxProcs))

	var wg sync.WaitGroup

	var minDividend uint32 = uint32(*flagDividendStart)
	var maxDividend uint32 = uint32(*flagDividendEnd) // fully tested for 0xffffff, may take a long time

	var minDivisor uint32 = uint32(*flagDivisorStart) // min is 2 because negative number and zero don't make sense, and 1 means every packet (no selection)
	var maxDivisor uint32 = uint32(*flagDivisorEnd)   // tested until 100000

	if *flagOutputAll {
		flag.Set("outputalldivisors", "true")
		flag.Set("outputalldifferences", "true")
	}

	var aDivisor uint32
	for aDivisor = minDivisor; aDivisor <= maxDivisor; aDivisor++ {
		go func(divisor uint32) {
			wg.Add(1)
			defer wg.Done()

			var dividend uint32
			var d Dividable

			// Check if divisor is a power of 2
			if divisor&(divisor-1) == 0 {
				if *flagIgnorePow2 {
					return
				} else {
					if *flagUsePow2 {
						d = NewDivisionPow2(divisor)
					} else {
						d = NewZeroremainder(divisor)
					}
				}
			} else {
				d = NewZeroremainder(divisor)
			}

			var zeroremainderCount, exactCount, zeroremainderWrong int
			for dividend = minDividend; dividend <= maxDividend; dividend++ {
				zeroremainder := d.IsRestlessDividable(dividend)
				exact := (dividend%divisor == 0)
				if *flagOutputAll {
					fmt.Println("CALC Dividend:", dividend, "Divisor:", divisor, "Zeroremainder:", d.IsRestlessDividable(dividend), "Exact:", (dividend%divisor == 0), "Correct:", zeroremainder == exact)
				}
				if zeroremainder != exact {
					zeroremainderWrong++
					if *flagOutputAllDifferences {
						fmt.Println("DIFF Dividend:", dividend, "Divisor:", divisor, "Zeroremainder:", d.IsRestlessDividable(dividend), "Exact:", (dividend%divisor == 0), "Reciprocal:", d.GetReciprocal())
					}
				}
				zeroremainderCount += Btoi(zeroremainder)
				exactCount += Btoi(exact)
			}
			if Round(1.0/float64(exactCount)*float64(zeroremainderCount), 4) != 1.0 || *flagOutputAllDivisors {
				fmt.Println("SUMMARY Divisor:", divisor, "Zeroremainder counts:", zeroremainderCount, "Modulo counts:", exactCount, "Difference (%):", 100.0-Round(100.0/float64(exactCount)*float64(zeroremainderCount), 4), "Differences between Zeroremainder and exact:", zeroremainderWrong, "Reciprocal:", d.GetReciprocal())
			}
		}(aDivisor)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
}
