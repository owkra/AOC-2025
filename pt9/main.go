package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type FreshRange struct {
	Start int
	End   int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("open input: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan input: %v", err)
	}

	var ranges []FreshRange
	for i := 0; i < len(lines) && i < 174; i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			log.Fatalf("invalid range on line %d: %q", i+1, line)
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("parse start on line %d: %v", i+1, err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("parse end on line %d: %v", i+1, err)
		}
		ranges = append(ranges, FreshRange{Start: start, End: end})
	}

	var ingredients []int
	for i := 174; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("parse ingredient id on input line %d: %v", i+1, err)
		}
		ingredients = append(ingredients, id)
	}

	res := FreshIngredients(ingredients, ranges)
	fmt.Printf("Fresh ingredients: %v", res)
}

func FreshIngredients(ingredientIds []int, freshRanges []FreshRange) int {
	sum := 0
	for _, ingredientID := range ingredientIds {
		for _, freshRange := range freshRanges {
			if ingredientID >= freshRange.Start && ingredientID <= freshRange.End {
				sum++
				break
			}
		}
	}
	return sum
}
