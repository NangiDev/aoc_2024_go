package day02

import (
	"AoC-2024/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day02_1() {

	count := 0

	for _, data := range strings.Split(utils.GetData(2, utils.Test), "\n") {
		is_safe := true
		data_arr := strings.Split(data, " ")

		int_arr := make([]int, len(data_arr))

		for i, str := range data_arr {
			num, _ := strconv.Atoi(str)
			int_arr[i] = num
		}

		ordered := func() bool {
			if slices.IsSorted(int_arr) {
				return true
			}
			reversed := make([]int, len(int_arr))
			copy(reversed, int_arr)
			slices.Reverse(reversed)
			return slices.IsSorted(reversed)
		}

		is_ordered := ordered()
		for i := 0; i < len(int_arr)-1; i++ {
			if !is_ordered {
				is_safe = false
				break
			}

			left := int_arr[i]
			right := int_arr[i+1]

			abs_diff := math.Abs((float64)(left - right))
			good_inc := abs_diff >= 1 && abs_diff <= 3

			if !(good_inc) {
				is_safe = false
				break
			}
		}

		if is_safe {
			count++
		}
	}

	println(count)
}
