package main

// mod
type DivisionMod struct {
	divisor uint64
}

func NewDivisionMod(divisor uint64) Dividable {
	divisionMod := &DivisionMod{divisor}
	return divisionMod
}

func (divisionMod *DivisionMod) IsRestlessDividable(dividend uint64) bool {
	return dividend%divisionMod.divisor == 0
}

func (divisionMod *DivisionMod) GetReciprocal() uint64 {
	return 0
}
