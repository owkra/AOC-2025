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

	res := FreshIngredientIds(ranges)
	fmt.Printf("Fresh ingredients: %v", res)
}

func FreshIngredientIds(freshRanges []FreshRange) int {
	sum := 0
	freshRanges = freshRangeConsolidation(freshRanges)
	for _, r := range freshRanges {
		sum += r.End - r.Start + 1
	}
	return sum
}

func freshRangeConsolidation(freshRanges []FreshRange) []FreshRange {
	for leftIdx, left := range freshRanges {
		for rightIdx, right := range freshRanges {
			if leftIdx == rightIdx {
				continue
			}
			if left.Start > right.End || right.Start > left.End {
				// no overlap
				continue
			} else if left.Start <= right.Start && left.End >= right.End {
				// full overlap
				freshRanges = append(freshRanges[:rightIdx], freshRanges[rightIdx+1:]...)
				return freshRangeConsolidation(freshRanges)
			} else if right.Start <= left.Start && right.End >= left.End {
				// full overlap
				freshRanges = append(freshRanges[:leftIdx], freshRanges[leftIdx+1:]...)
				return freshRangeConsolidation(freshRanges)
			} else {
				// partial overlap
				newStart := left.Start
				if right.Start < left.Start {
					newStart = right.Start
				}

				newEnd := left.End
				if right.End > left.End {
					newEnd = right.End
				}

				freshRanges[leftIdx] = FreshRange{Start: newStart, End: newEnd}
				freshRanges = append(freshRanges[:rightIdx], freshRanges[rightIdx+1:]...)
				return freshRangeConsolidation(freshRanges)
			}
		}
	}
	return freshRanges
}
