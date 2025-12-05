package main

import "testing"

func TestFreshIngredients_SingleRow(t *testing.T) {
	freshRanges := []FreshRange{
		{Start: 3, End: 5},
		{Start: 10, End: 14},
		{Start: 16, End: 20},
		{Start: 12, End: 18},
	}

	ingredients := []int{1, 5, 8, 11, 17, 32}

	want := 3

	got := FreshIngredients(ingredients, freshRanges)
	if got != want {
		t.Errorf("CheckAccessiblePaper() = %d; want %d", got, want)
	}
}
