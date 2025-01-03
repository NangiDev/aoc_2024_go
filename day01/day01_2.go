package day01

import (
	"AoC-2024/utils"
	"strconv"
	"strings"
)

func Day01_2() {
	var left []int
	var right []int

	var dict = make(map[int]int)

	for _, e := range strings.Split(utils.GetData(1, utils.Real), "\n") {
		chars := strings.Split(e, " ")

		leftNum, _ := strconv.Atoi(chars[0])
		rightNum, _ := strconv.Atoi(chars[3])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	for _, l := range left {
		count := 0
		for _, r := range right {
			if l == r {
				count++
			}
		}
		dict[l] = count
	}

	count := 0
	for _, v := range left {
		count += v * dict[v]
	}

	utils.AssertEqual(24941624, count)
	println(count)
}
