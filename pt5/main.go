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

	fmt.Printf("Battery bank: %v", bank)
}

type BatteryBank [][]int
