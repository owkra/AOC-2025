package main

import (
	"testing"
)

// Test template for pt5 — tabular tests and a single focused test.
// See implementation in [`pt5/main.go`](pt5/main.go:1).

func TestJoltageCalculation(t *testing.T) {
	tests := []struct {
		name string
		bank BatteryBank
		want int
	}{
		{
			name: "placeholder",
			bank: BatteryBank{
				// TODO: add rows, e.g. []int{1,2,3}
			},
			want: 0, // TODO: set expected value
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := JoltageCalculation(tc.bank)
			if got != tc.want {
				t.Errorf("JoltageCalculation(%v) = %d; want %d", tc.bank, got, tc.want)
			}
		})
	}
}

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
