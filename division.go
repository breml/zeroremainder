// type Dividable and two implementations, Division and DivisionPow2
package main

import (
	"math"
)

type Dividable interface {
	IsRestlessDividable(uint64) bool
	GetReciprocal() uint64
}

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

// Uint64
type ZeroremainderUint64 struct {
	divisor    uint64
	reciprocal uint64
}

func NewZeroremainderUint64(divisor uint64) Dividable {
	zeroremainder := &ZeroremainderUint64{divisor, math.MaxUint64/divisor + 1}
	return zeroremainder
}

func (zeroremainder *ZeroremainderUint64) IsRestlessDividable(dividend uint64) bool {
	if dividend == 0 {
		return true
	}
	if dividend < zeroremainder.divisor {
		return false
	}
	return (dividend * zeroremainder.reciprocal) < zeroremainder.reciprocal
}

func (zeroremainder *ZeroremainderUint64) GetReciprocal() uint64 {
	return zeroremainder.reciprocal
}

// Pow2
type DivisionPow2 struct {
	divisor    uint64
	reciprocal uint64
}

func NewDivisionPow2(divisor uint64) Dividable {
	divisionPow2 := &DivisionPow2{divisor, divisor - 1}
	return divisionPow2
}

func (divisionPow2 *DivisionPow2) IsRestlessDividable(dividend uint64) bool {
	if dividend == 0 {
		return true
	}
	if dividend < divisionPow2.divisor {
		return false
	}
	return (dividend & divisionPow2.reciprocal) == 0
}

func (divisionPow2 *DivisionPow2) GetReciprocal() uint64 {
	return divisionPow2.reciprocal
}
