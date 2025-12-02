package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rotation struct {
	Right bool
	Steps int
}

func loadRotations(path string) ([]Rotation, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []Rotation
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		dir := rune(line[0]) == 'R'
		n, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			return nil, fmt.Errorf("invalid rotation %q: %w", line, err)
		}
		out = append(out, Rotation{Right: dir, Steps: n})
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func main() {
	const input = "input.txt"
	rotations, err := loadRotations(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load rotations: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Final code: %d\n", CalculateCode(rotations))
}

func CalculateCode(rotations []Rotation) int {
	zeroCount := 0
	lockVal := 50
	for _, rotation := range rotations {
		if rotation.Steps > 100 {
			zeroCount += rotation.Steps / 100
			rotation.Steps = rotation.Steps % 100
		}

		alreadyZero := lockVal == 0
		if rotation.Right {
			// going positive direction
			lockVal += rotation.Steps
			if lockVal >= 100 {
				lockVal -= 100
				if !alreadyZero {
					zeroCount++
				}
			}
		} else {
			// going negative direction
			lockVal -= rotation.Steps
			if lockVal < 0 {
				lockVal += 100
				if !alreadyZero {
					zeroCount++
				}
			} else if lockVal == 0 {
				if !alreadyZero {
					zeroCount++
				}
			}
		}
	}
	return zeroCount
}
