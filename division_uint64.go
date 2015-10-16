package main

import (
	"math"
)

// Uint64
type ZeroremainderUint64 struct {
	divisor    uint64
	reciprocal uint64
}

func NewZeroremainderUint64(divisor uint64) Dividable {
	zeroremainder := &ZeroremainderUint64{uint64(divisor), math.MaxUint64/uint64(divisor) + 1}
	return zeroremainder
}

func (zeroremainder *ZeroremainderUint64) IsRestlessDividable(in_dividend uint64) bool {
	dividend := uint64(in_dividend)
	if dividend == 0 {
		return true
	}
	if dividend < zeroremainder.divisor {
		return false
	}
	return (dividend * zeroremainder.reciprocal) < zeroremainder.reciprocal
}

func (zeroremainder *ZeroremainderUint64) GetReciprocal() uint64 {
	return uint64(zeroremainder.reciprocal)
}
