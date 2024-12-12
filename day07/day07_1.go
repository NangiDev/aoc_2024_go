package day07

import (
	"AoC-2024/utils"
	"strconv"
	"strings"
)

func Day07_1() {
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

	utils.AssertEqual(3598800864292, count)
	println(count)
}

func calc(goal int, levels [][]int, v []int, ops []func(a int, b int) int) bool {
	for len(levels) > 0 && len(v) > 0 {
		level := levels[0]
		levels = levels[1:]
		b := v[0]
		v = v[1:]

		nLevel := []int{}
		for _, a := range level {
			for _, op := range ops {
				n := op(a, b)
				if n == goal {
					return true
				} else if n < goal {
					nLevel = append(nLevel, n)
				}
			}
		}
		levels = append(levels, nLevel)
	}
	return false
}

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}
