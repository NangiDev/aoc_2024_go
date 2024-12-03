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

	total := 0

	data := utils.GetData(1, utils.Test)
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
		total += int(math.Abs((float64)(l - r)))
	}

	println(total)
}
