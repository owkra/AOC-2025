package main

import (
	"testing"
)

func TestJoltageCalculation_SingleRow(t *testing.T) {
	// Single-test template (useful for focused unit test)
	bank := BatteryBank{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
	}
	want := 98 // TODO: expected result for the single row
	got := JoltageCalculation(bank)
	if got != want {
		t.Errorf("JoltageCalculation() = %d; want %d", got, want)
	}
}
