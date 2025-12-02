package main

import (
	"testing"
)

func TestCalculateCode(t *testing.T) {
	tests := []struct {
		name      string
		rotations []Rotation
		want      int
	}{
		{
			name: "basic",
			rotations: []Rotation{
				{Right: true, Steps: 150},
				{Right: false, Steps: 75},
			},
			want: 2,
		},
		{
			name: "puzzle-example",
			rotations: []Rotation{
				{Right: false, Steps: 68},
				{Right: false, Steps: 30},
				{Right: true, Steps: 48},
				{Right: false, Steps: 5},
				{Right: true, Steps: 60},
				{Right: false, Steps: 55},
				{Right: false, Steps: 1},
				{Right: false, Steps: 99},
				{Right: true, Steps: 14},
				{Right: false, Steps: 82},
			},
			want: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CalculateCode(tc.rotations)

			if got != tc.want {
				t.Errorf("CalculateCode() = %d; want %d", got, tc.want)
			}
		})
	}
}
