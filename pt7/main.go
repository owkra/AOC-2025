package main

import (
	"bufio"
	"fmt"
	"os"
)

type PaperGrid [][]rune

func main() {
	grid, err := ParsePaperGrid("input.txt")
	if err != nil {
		fmt.Printf("failed to parse paper grid: %v\n", err)
		return
	}

	fmt.Printf("Accessible Paper Count: %v", CheckAccessiblePaper(grid))
}

func CheckAccessiblePaper(grid PaperGrid) int {
	accessibleCount := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '.' {
				continue
			}

			left := x > 0
			right := x < len(grid)-1
			up := y > 0
			down := y < len(grid[x])-1

			surroundingPapers := 0
			if left {
				if grid[y][x-1] == '@' {
					surroundingPapers++
				}
			}
			if right {
				if grid[y][x+1] == '@' {
					surroundingPapers++
				}
			}
			if up {
				if grid[y-1][x] == '@' {
					surroundingPapers++
				}
			}
			if down {
				if grid[y+1][x] == '@' {
					surroundingPapers++
				}
			}
			if up && left {
				if grid[y-1][x-1] == '@' {
					surroundingPapers++
				}
			}
			if up && right {
				if grid[y-1][x+1] == '@' {
					surroundingPapers++
				}
			}
			if down && left {
				if grid[y+1][x-1] == '@' {
					surroundingPapers++
				}
			}
			if down && right {
				if grid[y+1][x+1] == '@' {
					surroundingPapers++
				}
			}

			if surroundingPapers < 4 {
				accessibleCount++
			}
		}
	}
	return accessibleCount
}

func ParsePaperGrid(path string) (PaperGrid, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make(PaperGrid, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// skip empty lines
		if len(line) == 0 {
			continue
		}
		row := []rune(line)
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}
