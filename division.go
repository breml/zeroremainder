// type Dividable and two implementations, Division and DivisionPow2
package main

import (
	"math"
)

type Dividable interface {
	IsRestlessDividable(uint32) bool
	GetReciprocal() uint32
}

type Zeroremainder struct {
	divisor    uint32
	reciprocal uint32
}

func NewZeroremainder(divisor uint32) Dividable {
	zeroremainder := &Zeroremainder{divisor, math.MaxUint32/divisor + 1}
	return zeroremainder
}

func (zeroremainder *Zeroremainder) IsRestlessDividable(dividend uint32) bool {
	if dividend == 0 {
		return true
	}
	if dividend < zeroremainder.divisor {
		return false
	}
	return (dividend * zeroremainder.reciprocal) < zeroremainder.reciprocal
}

func (zeroremainder *Zeroremainder) GetReciprocal() uint32 {
	return zeroremainder.reciprocal
}

type DivisionPow2 struct {
	divisor    uint32
	reciprocal uint32
}

func NewDivisionPow2(divisor uint32) Dividable {
	divisionPow2 := &DivisionPow2{divisor, divisor - 1}
	return divisionPow2
}

func (divisionPow2 *DivisionPow2) IsRestlessDividable(dividend uint32) bool {
	if dividend == 0 {
		return true
	}
	if dividend < divisionPow2.divisor {
		return false
	}
	return (dividend & divisionPow2.reciprocal) == 0
}

func (divisionPow2 *DivisionPow2) GetReciprocal() uint32 {
	return divisionPow2.reciprocal
}
