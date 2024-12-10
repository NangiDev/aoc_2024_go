package day01

import (
	"AoC-2024/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day01_1() {
	var left []int
	var right []int

	count := 0

	data := utils.GetData(1, utils.Real)
	for _, e := range strings.Split(data, "\n") {
		chars := strings.Split(e, " ")

		leftNum, _ := strconv.Atoi(chars[0])
		rightNum, _ := strconv.Atoi(chars[3])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	slices.Sort(left)
	slices.Sort(right)

	for i, l := range left {
		r := right[i]
		count += int(math.Abs((float64)(l - r)))
	}

	utils.AssertEqual(2086478, count)
	println(count)
}
