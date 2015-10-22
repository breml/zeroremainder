package main

import (
	"testing"
)

func BenchmarkDivisionMod(b *testing.B) {
	d := NewDivisionMod(1000)
	// run the dividable check function b.N times
	for n := 0; n < b.N; n++ {
		d.IsRestlessDividable(uint64(n))
	}
}

func BenchmarkDivisionPow2(b *testing.B) {
	d := NewDivisionPow2(1024)
	// run the dividable check function b.N times
	for n := 0; n < b.N; n++ {
		d.IsRestlessDividable(uint64(n))
	}
}

func BenchmarkZeroremainderUint32(b *testing.B) {
	d := NewZeroremainderUint32(1000)
	// run the dividable check function b.N times
	for n := 0; n < b.N; n++ {
		d.IsRestlessDividable(uint64(n))
	}
}

func BenchmarkZeroremainderUint64(b *testing.B) {
	d := NewZeroremainderUint64(1000)
	// run the dividable check function b.N times
	for n := 0; n < b.N; n++ {
		d.IsRestlessDividable(uint64(n))
	}
}
