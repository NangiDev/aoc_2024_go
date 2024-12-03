package day02

import (
	"AoC-2024/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day02_2() {

	var arr [][]int
	for _, data := range strings.Split(utils.GetData(2, utils.Test), "\n") {
		data_arr := strings.Split(data, " ")

		int_arr := make([]int, len(data_arr))

		for i, str := range data_arr {
			num, _ := strconv.Atoi(str)
			int_arr[i] = num
		}
		arr = append(arr, int_arr)
	}

	var potentially_unsafe [][]int
	count := 0
	for _, d := range arr {
		if isSafe(d) {
			count++
		} else {
			potentially_unsafe = append(potentially_unsafe, d)
		}
	}

	for _, d := range potentially_unsafe {
		for i := range len(d) {
			window := make([]int, len(d))
			copy(window, d)
			window = append(window[:i], window[i+1:]...)

			if isSafe(window) {
				count++
				break
			}
		}
	}

	println(count)
}

func isOrdered(intArr []int) bool {
	if slices.IsSorted(intArr) {
		return true
	}
	reversed := make([]int, len(intArr))
	copy(reversed, intArr)
	slices.Reverse(reversed)
	return slices.IsSorted(reversed)
}

func isSafe(intArr []int) bool {
	if !isOrdered(intArr) {
		return false
	}

	for i := 0; i < len(intArr)-1; i++ {
		absDiff := math.Abs(float64(intArr[i] - intArr[i+1]))
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return true
}
