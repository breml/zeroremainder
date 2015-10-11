// This file includes some helper functions
package main

import (
	"math"
)

// Convert boolean to int where true = 1 and false = 0
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// round a float to the next whole number
func round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}

// Round a float to given decimal
func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return round(f*shift) / shift
}
