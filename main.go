package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	flagDividendStart     = flag.Int("dividendstart", 0, "Start of range of dividends")
	flagDividendEnd       = flag.Int("dividendend", 0xffff, "End of range of dividends")
	flagDivisorStart      = flag.Int("divisorstart", 2, "Start of range of divisors")
	flagDivisorEnd        = flag.Int("divisorend", 1000, "End of range of divisors")
	flagMaxProcs          = flag.Int("maxprocs", 0, "Value for GOMAXPROCS (value may be reduced, respecting default GOMAXPROCS and number of CPUs)")
	flagOutputAllDivisors = flag.Bool("outputalldivisors", false, "Output result for all divisors")
	flagOutputDifferences = flag.Bool("outputdifferences", false, "Output every dividend / divisor combination which provides a flase result")
	flagOutputCalc        = flag.Bool("outputcalc", false, "Output every calculation; WARNING: very verbose, use only with small ranges for dividend and divisor, maxprocs should be set to 1")
	flagIgnorePow2        = flag.Bool("ignorepow2", true, "for power of 2 divisors the solution is trivial, therefore these divisors may be ignored")
	flagUsePow2           = flag.Bool("usepow2", true, "use bitwise and for power of 2 divisors")
	flagHelp              = flag.Bool("help", false, "Print help")
	flagNumType           = flag.String("numtype", "uint32", "numeric type, used for the calculations (one of: uint8, uint16, uint32, uint64)")
)

func main() {
	flag.Parse()

	if *flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	numType := strings.ToLower(*flagNumType)
	switch numType {
	case "uint8", "uint16", "uint32", "uint64":
	default:
		fmt.Println("numtype must be one of: uint8, uint16, uint32, uint64")
		os.Exit(1)
	}

	runtime.GOMAXPROCS(MaxParallelism(*flagMaxProcs))

	var wg sync.WaitGroup

	var minDividend uint64 = uint64(*flagDividendStart)
	var maxDividend uint64 = uint64(*flagDividendEnd) // fully tested for 0xffffff, may take a long time

	var minDivisor uint64 = uint64(*flagDivisorStart) // min is 2 because negative number and zero don't make sense, and 1 means every packet (no selection)
	var maxDivisor uint64 = uint64(*flagDivisorEnd)   // tested until 100000

	var aDivisor uint64
	for aDivisor = minDivisor; aDivisor <= maxDivisor; aDivisor++ {
		go func(divisor uint64) {
			wg.Add(1)
			defer wg.Done()

			var dividend uint64
			var d Dividable

			// Check if divisor is a power of 2
			if divisor&(divisor-1) == 0 {
				if *flagIgnorePow2 {
					return
				}
				if *flagUsePow2 {
					d = NewDivisionPow2(divisor)
				}
			}

			if d == nil {
				switch numType {
				case "uint8":
					d = NewZeroremainderUint8(divisor)
				case "uint16":
					d = NewZeroremainderUint16(divisor)
				case "uint32":
					d = NewZeroremainderUint32(divisor)
				case "uint64":
					d = NewZeroremainderUint64(divisor)
				}
			}
			e := NewDivisionMod(divisor)

			var zeroremainderCount, exactCount, zeroremainderWrong int
			for dividend = minDividend; dividend <= maxDividend; dividend++ {
				zeroremainder := d.IsRestlessDividable(dividend)
				exact := e.IsRestlessDividable(dividend)
				if *flagOutputCalc {
					fmt.Printf("CALC Dividend: %9d Divisor: %9d Zeroremainder: %5t Exact: %5t Correct: %5t\n", dividend, divisor, d.IsRestlessDividable(dividend), dividend%divisor == 0, zeroremainder == exact)
				}
				if zeroremainder != exact {
					zeroremainderWrong++
					if *flagOutputDifferences {
						fmt.Printf("DIFF Dividend: %9d Divisor: %9d Zeroremainder: %5t Exact: %5t Reciprocal: %9d\n", dividend, divisor, d.IsRestlessDividable(dividend), dividend%divisor == 0, d.GetReciprocal())
					}
				}
				zeroremainderCount += Btoi(zeroremainder)
				exactCount += Btoi(exact)
			}
			if Round(1.0/float64(exactCount)*float64(zeroremainderCount), 4) != 1.0 || *flagOutputAllDivisors {
				fmt.Printf("SUMM Divisor: %9d Zeroremainder counts: %9d Modulo counts: %9d Difference (%%): %.6f Differences between Zeroremainder and exact: %9d Reciprocal: %9d\n", divisor, zeroremainderCount, exactCount, 100.0-Round(100.0/float64(exactCount)*float64(zeroremainderCount), 4), zeroremainderWrong, d.GetReciprocal())
			}
		}(aDivisor)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
}
