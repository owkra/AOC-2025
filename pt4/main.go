package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IDRange struct {
	Min int
	Max int
}

// parseRanges takes a CSV line of ranges in the form "##-##,##-##,..." and
// returns a slice of IDRange. It skips empty entries and returns an error if
// any range is malformed.
func parseRanges(s string) ([]IDRange, error) {
	var out []IDRange
	parts := strings.Split(s, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		both := strings.Split(p, "-")
		if len(both) != 2 {
			return nil, fmt.Errorf("invalid range %q", p)
		}
		min, err := strconv.Atoi(strings.TrimSpace(both[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid min in %q: %w", p, err)
		}
		max, err := strconv.Atoi(strings.TrimSpace(both[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid max in %q: %w", p, err)
		}
		out = append(out, IDRange{Min: min, Max: max})
	}
	return out, nil
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		return
	}
	line := strings.TrimSpace(string(data))
	ranges, err := parseRanges(line)
	if err != nil {
		fmt.Printf("failed to parse ranges: %v\n", err)
		return
	}

	invalidIDs := InvalidIDs(ranges)
	sum := 0
	for _, id := range invalidIDs {
		sum += id
	}
	fmt.Printf("Sum of invalid IDs: %d\n", sum)
}

func InvalidIDs(idRanges []IDRange) []int {
	ids := make([]int, 0)
	for _, idRange := range idRanges {
		for id := idRange.Min; id <= idRange.Max; id++ {
			strId := strconv.Itoa(id)

			matchStr := ""
			isInvalid := false
			for matchLen := 1; matchLen <= len(strId)/2; matchLen++ {
				matchStr = string(strId[:matchLen])

				strValid := false
				for strCheck := 0; strCheck < len(strId); strCheck += matchLen {
					isLongEnough := strCheck+matchLen <= len(strId)
					if !isLongEnough {
						strValid = true
						break
					}

					if strId[strCheck:strCheck+matchLen] != matchStr {
						strValid = true
						break
					}
				}

				if !strValid {
					isInvalid = true
					break
				}

			}

			if isInvalid {
				ids = append(ids, id)
			}
		}
	}
	return ids
}
