package main

import (
	"bufio"
	"fmt"
	"os"
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
		// scan through the row once, pick the greatest that furthest to the left and not at the end

		tensIdx := 0
		for i := 0; i < len(row)-1; i++ {
			val := row[i]
			if val > row[tensIdx] {
				tensIdx = i
			}
		}

		onesRow := row[tensIdx+1:]
		onesIdx := 0
		for i := 0; i < len(onesRow); i++ {
			val := onesRow[i]
			if val >= onesRow[onesIdx] {
				onesIdx = i
			}
		}

		totalJoltage += row[tensIdx]*10 + onesRow[onesIdx]
	}
	return totalJoltage
}
