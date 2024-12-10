package day06

import (
	"AoC-2024/utils"
	"slices"
	"strings"
)

func initialize(start *Pos, board *[][]string, input *[]string) {
	for y, line := range *input {
		if strings.Contains(line, "^") {
			start.x = strings.Index(line, "^")
			start.y = y
		}
		row := strings.Split(line, "")
		*board = append(*board, row)
	}
}

const maxSteps = 100000

func walk(start *Pos, dir Pos, board [][]string) ([]Pos, int32) {
	path := []Pos{}
	flow := make(map[Pos][]Pos)
	tile := getValid(*start, board)
	steps := 0
	for tile != nil {
		if slices.Contains(flow[*start], dir) ||
			steps >= maxSteps {
			return path, 1
		}

		next := Pos{
			start.x + dir.x,
			start.y + dir.y,
		}
		nextTile := getValid(next, board)
		for nextTile != nil && *nextTile == "#" {
			dir = directions[dir]
			next = Pos{
				start.x + dir.x,
				start.y + dir.y,
			}
			nextTile = getValid(next, board)
		}

		path = append(path, *start)
		flow[*start] = append(flow[*start], dir)

		start.x += dir.x
		start.y += dir.y
		tile = getValid(*start, board)
		steps++
	}

	// for _, e := range board {
	// 	fmt.Println(e)
	// }
	// fmt.Println()

	return path, 0
}

func Day06_2() {
	input := strings.Split(utils.GetData(6, utils.Real), "\n")
	start := Pos{0, 0}
	board := [][]string{}
	initialize(&start, &board, &input)

	upDir := Pos{0, -1}
	path, _ := walk(&start, upDir, board)

	var count int32 = 0

	for i := 0; i < len(path)-1; i++ {
		cur := path[i]
		next := path[i+1]
		dir := Pos{
			next.x - cur.x,
			next.y - cur.y,
		}

		newBoard := make([][]string, len(board))
		for i := range board {
			newBoard[i] = make([]string, len((board)[i]))
			copy(newBoard[i], (board)[i])
		}

		// X and Y is mixed.
		newBoard[next.y][next.x] = "#"

		_, loop := walk(&cur, dir, newBoard)
		count += loop
	}

	// for _, e := range board {
	// 	fmt.Println(e)
	// }
	// fmt.Println()

	println(count)
}
