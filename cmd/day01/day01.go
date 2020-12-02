package main

import (
	"fmt"
	"os"

	"github.com/Kreonn/AoC2020/internal/tools"
)

func getPartOne(data []int) int {
	for _, first := range data {
		for _, second := range data {
			if first+second == 2020 {
				return first * second
			}
		}
	}

	return 0
}

func getPartTwo(data []int) int {
	for _, first := range data {
		for _, second := range data {
			for _, third := range data {
				if first+second+third == 2020 {
					return first * second * third
				}
			}
		}
	}

	return 0
}

func main() {
	data, err := tools.GetData("day01.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	intData, err := tools.Str2Int(data.Lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Part 1
	fmt.Printf("Result for part 1 : %d\n", getPartOne(intData))

	// Part 2
	fmt.Printf("Result for part 2 : %d\n", getPartTwo(intData))
}
