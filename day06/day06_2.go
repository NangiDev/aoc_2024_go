package day06

import (
	"AoC-2024/utils"
	"slices"
	"strings"
)

func Day06_2() {
	input := strings.Split(utils.GetData(6, utils.Test), "\n")

	guard := Pos{0, 0}
	board := [][]string{}
	for y, line := range input {
		if x := strings.Index(line, "^"); x > 0 {
			guard = Pos{x, y}
		}

		line = strings.ReplaceAll(line, ".", " ")

		board = append(board, strings.Split(line, ""))
	}

	flow := make(map[Pos][]Pos)
	dir := Pos{0, -1}
	count := 0
	isValid, tile := getValid(guard, board)
	for isValid {
		*tile = "+"
		flow[guard] = append(flow[guard], dir)

		right := Pos{
			guard.x + directions[dir].x,
			guard.y + directions[dir].y,
		}
		validRight, rightTile := getValid(right, board)
		for validRight && *rightTile != "#" {
			if slices.Contains(flow[right], directions[dir]) {
				count++
				break
			}
			right.x += directions[dir].x
			right.y += directions[dir].y
			validRight, rightTile = getValid(right, board)
		}

		next := Pos{
			guard.x + dir.x,
			guard.y + dir.y,
		}

		validNext, nextTile := getValid(next, board)
		if validNext && *nextTile == "#" {
			dir = directions[dir]
		}
		guard.x += dir.x
		guard.y += dir.y

		isValid, tile = getValid(guard, board)
	}

	// for _, r := range board {
	// 	fmt.Println(r)
	// }
	// fmt.Println()
	// println(count)
}
