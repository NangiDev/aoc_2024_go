package day04

import (
	"AoC-2024/utils"
	"strings"
)

var two_dir = []Pos{
	{-1, -1},
	{1, -1},
}

func Day04_2() {
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
			if table[x][y] == "A" {
				pos := Pos{x, y}
				starts = append(starts, pos)
			}
		}
	}

	count := 0
	// Traverse all XMAS starting from X in all Directions
	for _, s := range starts {
		is_x := false
		for _, dir := range two_dir {
			dx := s.x + dir.x
			dy := s.y + dir.y
			neg_dx := s.x - dir.x
			neg_dy := s.y - dir.y

			cur, valid_cur := getValid(s.x, s.y, table)
			cur_q1, valid_q1 := getValid(dx, dy, table)
			cur_q2, valid_q2 := getValid(neg_dx, neg_dy, table)

			if valid_cur && valid_q1 && valid_q2 {
				if cur != "A" || cur_q1 == cur_q2 {
					is_x = false
					break
				} else if cur_q1 == "M" && cur_q2 == "S" {
					is_x = true
				} else if cur_q1 == "S" && cur_q2 == "M" {
					is_x = true
				} else {
					is_x = false
					break
				}
			}
		}
		if is_x {
			count++
		}
	}

	println(count)
}

func getValid(x int, y int, table [][]string) (string, bool) {
	width := len(table[0])
	height := len(table)

	// IN BOUNDS
	if (x >= 0 && x < width) && (y >= 0 && y < height) {
		return table[x][y], true
	}

	// OUT OF BOUNDS
	return "", false
}
