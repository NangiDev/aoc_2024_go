package day07

import (
	"AoC-2024/utils"
	"strconv"
	"strings"
)

func Day07_2() {
	input := strings.Split(utils.GetData(7, utils.Real), "\n")

	goals := make(map[int][]int)
	for _, e := range input {
		row := strings.Split(e, ": ")
		goal, _ := strconv.Atoi(row[0])
		numbers := []int{}
		for _, e := range strings.Split(row[1], " ") {
			n, _ := strconv.Atoi(e)
			numbers = append(numbers, n)
		}
		goals[goal] = numbers
	}

	var ops = []func(a int, b int) int{
		add,
		mul,
		concat,
	}

	count := 0
	for goal, v := range goals {
		levels := [][]int{}
		levels = append(levels, v[:1])
		v = v[1:]

		if calc(goal, levels, v, ops) {
			count += goal
		}

	}

	// Answer 340481958202782 is to high
	println(count)
}

func concat(a int, b int) int {
	multiplier := 1
	for multiplier <= b {
		multiplier *= 10
	}
	return a*multiplier + b
}
