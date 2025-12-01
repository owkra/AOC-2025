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

	zeroCount := 0
	lockVal := 50
	for _, rot := range rotations {
		if rot.Right {
			lockVal -= rot.Steps
			lockVal = lockVal % 100
		} else {
			lockVal += rot.Steps
			lockVal = lockVal % 100
		}

		if lockVal == 0 {
			zeroCount++
		}
	}

	fmt.Printf("Final code: %d\n", zeroCount)
}
