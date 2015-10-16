package main

import (
	"math"
)

// Uint16
type ZeroremainderUint16 struct {
	divisor    uint16
	reciprocal uint16
}

func NewZeroremainderUint16(divisor uint64) Dividable {
	zeroremainder := &ZeroremainderUint16{uint16(divisor), math.MaxUint16/uint16(divisor) + 1}
	return zeroremainder
}

func (zeroremainder *ZeroremainderUint16) IsRestlessDividable(in_dividend uint64) bool {
	dividend := uint16(in_dividend)
	if dividend == 0 {
		return true
	}
	if dividend < zeroremainder.divisor {
		return false
	}
	return (dividend * zeroremainder.reciprocal) < zeroremainder.reciprocal
}

func (zeroremainder *ZeroremainderUint16) GetReciprocal() uint64 {
	return uint64(zeroremainder.reciprocal)
}
