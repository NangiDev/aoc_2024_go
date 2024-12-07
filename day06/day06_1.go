package day06

import (
	"AoC-2024/utils"
	"strings"
)

type Pos struct {
	x int
	y int
}

var directions = map[Pos]Pos{
	{0, -1}: {1, 0},
	{1, 0}:  {0, 1},
	{0, 1}:  {-1, 0},
	{-1, 0}: {0, -1},
}

func Day06_1() {
	input := strings.Split(utils.GetData(6, utils.Real), "\n")

	guard := Pos{0, 0}
	dir := Pos{0, -1}
	board := [][]string{}
	for y, line := range input {
		if strings.Contains(line, "^") {
			guard.x = strings.Index(line, "^")
			guard.y = y
		}
		row := strings.Split(line, "")
		board = append(board, row)
	}

	count := 0
	isValid, tile := getValid(guard, board)
	for isValid && tile != nil {
		if *tile == "." || *tile == "^" {
			*tile = "X"
			count++
		}

		d := Pos{
			guard.x + dir.x,
			guard.y + dir.y,
		}
		_, tile = getValid(d, board)
		if tile != nil && *tile == "#" {
			dir = directions[dir]
		}
		guard.x += dir.x
		guard.y += dir.y
		isValid, tile = getValid(guard, board)
	}

	println(count)
}

func getValid(pos Pos, board [][]string) (bool, *string) {
	width := len(board[0])
	height := len(board)

	// IN BOUNDS
	if (pos.x >= 0 && pos.x < width) && (pos.y >= 0 && pos.y < height) {
		return true, &board[pos.y][pos.x]
	}

	// OUT OF BOUNDS
	return false, nil
}
