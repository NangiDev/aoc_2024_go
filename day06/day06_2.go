package day06

import (
	"AoC-2024/utils"
	"slices"
	"strings"
)

func Day06_2() {
	input := strings.Split(utils.GetData(6, utils.Real), "\n")
	start := Pos{0, 0}
	board := [][]string{}
	initialize(&start, &board, &input)

	count := 0
	walk(&start, &board, &count)

	// for _, e := range board {
	// 	fmt.Println(e)
	// }

	println(count)
}

func testLoop(start Pos, dir Pos, flow map[Pos][]Pos, board [][]string, count *int) {
	next := Pos{
		start.x + dir.x,
		start.y + dir.y,
	}
	nextTile := getValid(next, board)
	if nextTile != nil {
		*nextTile = "#"
		dir = directions[dir]
	}

	tile := getValid(start, board)
	for tile != nil {
		if slices.Contains(flow[start], dir) {
			*count++
			break
		}

		// Next step
		next := Pos{
			start.x + dir.x,
			start.y + dir.y,
		}

		// Check if should turn
		nextTile := getValid(next, board)
		if nextTile == nil {
			break
		} else if *nextTile == "#" {
			dir = directions[dir]
		}

		flow[start] = append(flow[start], dir)

		// Move to next
		start.x += dir.x
		start.y += dir.y
		tile = getValid(start, board)
	}
}

func walk(start *Pos, board *[][]string, count *int) {
	dir := Pos{0, -1}
	tile := getValid(*start, *board)
	flow := make(map[Pos][]Pos)

	for tile != nil {
		if slices.Contains(flow[*start], dir) {
			break
		}

		// Next step
		next := Pos{
			start.x + dir.x,
			start.y + dir.y,
		}

		// Check if should turn
		nextTile := getValid(next, *board)
		if nextTile == nil {
			break
		} else if *nextTile == "#" {
			dir = directions[dir]
		} else {
			// Copy Flow
			newFlow := make(map[Pos][]Pos, len(flow))
			for key, value := range flow {
				copiedSlice := make([]Pos, len(value))
				copy(copiedSlice, value)
				newFlow[key] = copiedSlice
			}

			// Copy board
			newBoard := make([][]string, len(*board))
			for i := range *board {
				newBoard[i] = make([]string, len((*board)[i]))
				copy(newBoard[i], (*board)[i])
			}

			testLoop(*start, dir, newFlow, newBoard, count)
		}

		flow[*start] = append(flow[*start], dir)

		// Move to next
		start.x += dir.x
		start.y += dir.y
		tile = getValid(*start, *board)
	}
}

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
