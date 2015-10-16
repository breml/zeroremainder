package main

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
