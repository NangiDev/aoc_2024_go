package day03

import (
	"fmt"
	"strconv"
	"strings"
)

var test_input_2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func Day03_2() {
	stack := test_input_2
	first_i := -1
	last_i := -1
	count := 0
	doMul := true
	for i, c := range stack {
		if string(c) == "(" {
			first_i = i
			last_i = -1
		} else if string(c) == ")" {
			last_i = i
		}

		if first_i >= 0 && last_i >= 0 {
			numbers := stack[first_i+1 : last_i]

			do := stack[max(first_i-2, 0):first_i]
			dont := stack[max(first_i-5, 0):first_i]

			if do == "do" {
				doMul = true
			}
			if dont == "don't" {
				doMul = false
			}

			op := stack[max(first_i-3, 0):first_i]
			if len(numbers) > 0 && op == "mul" && doMul {
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

	fmt.Println(count)
}
