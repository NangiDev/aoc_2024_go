package day04

import (
	"AoC-2024/utils"
	"strings"
)

type Pos struct {
	x int
	y int
}

var eight_dir = []Pos{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func Day04_1() {
	data := strings.Split(utils.GetData(4, utils.Real), "\n")

	// Convert to 2D table
	table := [][]string{}
	for _, d := range data {
		col := strings.Split(string(d), "")
		table = append(table, col)
	}

	// Find starting points
	starts := []Pos{}
	for y, col := range table {
		for x := range col {
			if table[x][y] == "X" {
				pos := Pos{x, y}
				starts = append(starts, pos)
			}
		}
	}

	count := 0
	width := len(table[0])
	height := len(table)

	// Traverse all XMAS starting from X in all Directions
	for _, s := range starts {
		for _, dir := range eight_dir {
			prev := "X"
			for i := 1; i <= 3; i++ {
				dx := s.x + dir.x*i
				dy := s.y + dir.y*i
				// Check if in bounds
				if (dx >= 0 && dx < width) && (dy >= 0 && dy < height) {
					cur := table[dx][dy]
					if (prev == "X") && (cur == "M") ||
						(prev == "M") && (cur == "A") {
						prev = cur
					} else if (prev == "A") && (cur == "S") {
						count++
					} else {
						break
					}
				}
				// OUT OF BOUNDS
			}
		}
	}

	utils.AssertEqual(2524, count)
	println(count)
}
