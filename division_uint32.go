package main

import (
	"math"
)

// Uint32
type ZeroremainderUint32 struct {
	divisor    uint32
	reciprocal uint32
}

func NewZeroremainderUint32(divisor uint64) Dividable {
	zeroremainder := &ZeroremainderUint32{uint32(divisor), math.MaxUint32/uint32(divisor) + 1}
	return zeroremainder
}

func (zeroremainder *ZeroremainderUint32) IsRestlessDividable(in_dividend uint64) bool {
	dividend := uint32(in_dividend)
	if dividend == 0 {
		return true
	}
	if dividend < zeroremainder.divisor {
		return false
	}
	return (dividend * zeroremainder.reciprocal) < zeroremainder.reciprocal
}

func (zeroremainder *ZeroremainderUint32) GetReciprocal() uint64 {
	return uint64(zeroremainder.reciprocal)
}
