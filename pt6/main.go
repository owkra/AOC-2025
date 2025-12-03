package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("failed to read input file: %v\n", err)
		return
	}
	defer file.Close()

	bank := make(BatteryBank, 0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		batteryRow := make([]int, 0, len(line))
		for _, digit := range line {
			batteryRow = append(batteryRow, int(digit-'0'))
		}
		bank = append(bank, batteryRow)
	}

	fmt.Printf("Joltage: %v", JoltageCalculation(bank))
}

type BatteryBank [][]int

func JoltageCalculation(bank BatteryBank) int {
	totalJoltage := 0
	for _, row := range bank {
		digitStore := ""
		prevIdx := -1
		for i := 0; i < 12; i++ {
			digit := 0

			maxIdx := 12 - i
			minIdx := prevIdx + 1
			for j := minIdx; j <= len(row)-maxIdx; j++ {
				val := row[j]
				if val > digit {
					digit = val
					prevIdx = j
				}
			}

			digitStore += fmt.Sprintf("%d", digit)
		}

		parsedVal, _ := strconv.Atoi(digitStore)

		totalJoltage += parsedVal
	}
	return totalJoltage
}
