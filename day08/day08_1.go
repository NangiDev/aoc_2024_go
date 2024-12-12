package day08

import (
	"AoC-2024/utils"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Antenna struct {
	pos  Pos
	freq string
}

func Day08_1() {
	input := strings.Split(utils.GetData(8, utils.Real), "\n")

	grid := [][]string{}
	antennas := []Antenna{}

	for y, i := range input {
		grid = append(grid, strings.Split(i, ""))
		for x, j := range strings.Split(i, "") {
			if j != "." {
				antennas = append(antennas, Antenna{Pos{x, y}, j})
			}
		}
	}

	width := len(input[0])
	height := len(input)
	unique := make(map[Pos]string)

	for i, a := range antennas {
		for _, b := range antennas[i+1:] {
			if a.freq == b.freq {
				dir := Pos{
					a.pos.x - b.pos.x,
					a.pos.y - b.pos.y,
				}
				negdir := Pos{
					-dir.x,
					-dir.y,
				}

				antiA := Pos{
					a.pos.x + dir.x,
					a.pos.y + dir.y,
				}
				antiB := Pos{
					b.pos.x + negdir.x,
					b.pos.y + negdir.y,
				}

				if antiA.x >= 0 && antiA.y >= 0 && antiA.x < width && antiA.y < height {
					grid[antiA.y][antiA.x] = "#"
					unique[antiA] = ""
				}

				if antiB.x >= 0 && antiB.y >= 0 && antiB.x < width && antiB.y < height {
					grid[antiB.y][antiB.x] = "#"
					unique[antiB] = ""
				}
			}
		}
	}

	utils.AssertEqual(426, len(unique))
	println(len(unique))
}
