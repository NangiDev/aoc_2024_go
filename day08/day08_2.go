package day08

import (
	"AoC-2024/utils"
	"strings"
)

func Day08_2() {
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
				antiA := Pos{
					a.pos.x,
					a.pos.y,
				}
				antiB := Pos{
					b.pos.x,
					b.pos.y,
				}

				dir := Pos{
					a.pos.x - b.pos.x,
					a.pos.y - b.pos.y,
				}
				negdir := Pos{
					-dir.x,
					-dir.y,
				}

				for (antiA.x >= 0 && antiA.x < width && antiA.y >= 0 && antiA.y < height) ||
					(antiB.x >= 0 && antiB.x < width && antiB.y >= 0 && antiB.y < height) {

					if antiA.x >= 0 && antiA.x < width && antiA.y >= 0 && antiA.y < height {
						grid[antiA.y][antiA.x] = "#"
						unique[antiA] = ""
					}

					if antiB.x >= 0 && antiB.x < width && antiB.y >= 0 && antiB.y < height {
						grid[antiB.y][antiB.x] = "#"
						unique[antiB] = ""
					}

					antiA = Pos{
						antiA.x + dir.x,
						antiA.y + dir.y,
					}
					antiB = Pos{
						antiB.x + negdir.x,
						antiB.y + negdir.y,
					}
				}
			}
		}
	}

	utils.AssertEqual(1359, len(unique))
	println(len(unique))
}
