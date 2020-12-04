package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Kreonn/AoC2020/internal/tools"
)

type Slope struct {
	right int
	down  int
}

func getTreesOnSlope(forest [][]string, slope Slope) int {
	total := 0
	rightCnt := 0
	for i := slope.down; i < len(forest); i += slope.down {
		rightCnt = (rightCnt + slope.right) % len(forest[i])
		if forest[i][rightCnt] == "#" {
			total++
		}
	}

	return total
}

func main() {
	data, err := tools.GetData("day03.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Populate forest
	var forest [][]string
	forest = make([][]string, len(data.Lines))
	for i := range forest {
		forest[i] = make([]string, len(data.Lines[0]))
		forest[i] = strings.Split(data.Lines[i], "")
	}

	slope1 := Slope{3, 1}
	total1 := getTreesOnSlope(forest, slope1)

	fmt.Printf("Result for part 1 : %d\n", total1)

	slope2 := Slope{1, 1}
	slope3 := Slope{5, 1}
	slope4 := Slope{7, 1}
	slope5 := Slope{1, 2}

	total2 := getTreesOnSlope(forest, slope2)
	total3 := getTreesOnSlope(forest, slope3)
	total4 := getTreesOnSlope(forest, slope4)
	total5 := getTreesOnSlope(forest, slope5)

	fmt.Printf("Result for part 2 : %d\n", total1*total2*total3*total4*total5)
}
