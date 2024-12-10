package day03

import (
	"AoC-2024/utils"
	"strconv"
	"strings"
)

func Day03_1() {
	stack := utils.GetData(3, utils.Real)

	first_i := -1
	last_i := -1
	count := 0
	for i, c := range stack {
		if string(c) == "(" {
			first_i = i
			last_i = -1
		} else if string(c) == ")" {
			last_i = i
		}

		if first_i >= 0 && last_i >= 0 {
			numbers := stack[first_i+1 : last_i]
			op := stack[max(first_i-3, 0):first_i]
			if len(numbers) > 0 && op == "mul" {
				num_split := strings.Split(numbers, ",")
				if len(num_split) == 2 {
					left, _ := strconv.Atoi(num_split[0])
					right, _ := strconv.Atoi(num_split[1])
					count += left * right
				}
			}
			first_i = -1
			last_i = -1
		}
	}

	utils.AssertEqual(175615763, count)
	println(count)
}
