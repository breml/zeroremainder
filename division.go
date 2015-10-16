package main

// type Dividable for Zeroremainder{8,16,31,64} and DivisionPow2
type Dividable interface {
	IsRestlessDividable(uint64) bool
	GetReciprocal() uint64
}
