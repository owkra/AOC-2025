package main

import (
	"testing"
)

func TestInvalidIDs(t *testing.T) {
	tests := []struct {
		name     string
		idRanges []IDRange
		want     int
	}{
		{
			name: "basic",
			idRanges: []IDRange{
				{Min: 11, Max: 22},
			},
			want: 33,
		},
		{
			name: "puzzle-example",
			idRanges: []IDRange{
				{Min: 11, Max: 22},
				{Min: 95, Max: 115},
				{Min: 998, Max: 1012},
				{Min: 1188511880, Max: 1188511890},
				{Min: 222220, Max: 222224},
				{Min: 1698522, Max: 1698528},
				{Min: 446443, Max: 446449},
				{Min: 38593856, Max: 38593862},
				{Min: 565653, Max: 565659},
				{Min: 824824821, Max: 824824827},
				{Min: 2121212118, Max: 2121212124},
			},
			want: 4174379265,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ids := InvalidIDs(tc.idRanges)

			got := 0
			for _, id := range ids {
				got += id
			}

			if got != tc.want {
				t.Errorf("InvalidIDs() = %d; want %d", got, tc.want)
			}
		})
	}
}

func TestInvalidIDsIndividual(t *testing.T) {
	idRange := []IDRange{
		{Min: 2121212118, Max: 2121212124},
	}
	want := 2121212121

	ids := InvalidIDs(idRange)

	got := 0
	for _, id := range ids {
		got += id
	}

	if got != want {
		t.Errorf("InvalidIDs() = %d; want %d", got, want)
	}
}
