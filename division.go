// type Dividable and two implementations, Division and DivisionPow2
package main

import (
	"math"
)

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
