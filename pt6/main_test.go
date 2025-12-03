package main

import (
	"testing"
)

func TestJoltageCalculation(t *testing.T) {
	tests := []struct {
		name string
		bank BatteryBank
		want int
	}{
		{
			name: "placeholder",
			bank: BatteryBank{
				[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
				[]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
				[]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
				[]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			},
			want: 3121910778619, // TODO: set expected value
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
	want := 987654321111
	got := JoltageCalculation(bank)
	if got != want {
		t.Errorf("JoltageCalculation() = %d; want %d", got, want)
	}
}
