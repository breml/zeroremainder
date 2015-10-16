package main

import (
	"math"
)

// Uint8
type ZeroremainderUint8 struct {
	divisor    uint8
	reciprocal uint8
}

func NewZeroremainderUint8(divisor uint64) Dividable {
	zeroremainder := &ZeroremainderUint8{uint8(divisor), math.MaxUint8/uint8(divisor) + 1}
	return zeroremainder
}

func (zeroremainder *ZeroremainderUint8) IsRestlessDividable(in_dividend uint64) bool {
	dividend := uint8(in_dividend)
	if dividend == 0 {
		return true
	}
	if dividend < zeroremainder.divisor {
		return false
	}
	return (dividend * zeroremainder.reciprocal) < zeroremainder.reciprocal
}

func (zeroremainder *ZeroremainderUint8) GetReciprocal() uint64 {
	return uint64(zeroremainder.reciprocal)
}
