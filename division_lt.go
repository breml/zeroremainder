package main

import (
	"math"
)

// lower than reciprocal
type DivisionLt struct {
	divisor    uint64
	reciprocal uint64
}

func NewDivisionLt(divisor uint64) Dividable {
	divisionLt := &DivisionLt{uint64(divisor), math.MaxUint64 / uint64(divisor)}
	return divisionLt
}

func (divisionLt *DivisionLt) IsRestlessDividable(dividend uint64) bool {
	return dividend < divisionLt.reciprocal
}

func (divisionLt *DivisionLt) GetReciprocal() uint64 {
	return divisionLt.reciprocal
}
