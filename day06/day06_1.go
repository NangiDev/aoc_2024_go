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
	tile := getValid(guard, board)
	for tile != nil {
		if *tile == "." || *tile == "^" {
			*tile = "X"
			count++
		}

		d := Pos{
			guard.x + dir.x,
			guard.y + dir.y,
		}
		tile = getValid(d, board)
		if tile != nil && *tile == "#" {
			dir = directions[dir]
		}
		guard.x += dir.x
		guard.y += dir.y
		tile = getValid(guard, board)
	}

	utils.AssertEqual(4665, count)
	println(count)
}

func getValid(pos Pos, board [][]string) *string {
	// Check boundaries
	if pos.x < 0 || pos.y < 0 || pos.y >= len(board) || pos.x >= len(board[pos.y]) {
		return nil
	}
	return &board[pos.y][pos.x]
}
