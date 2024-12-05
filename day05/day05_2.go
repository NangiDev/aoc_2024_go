package day05

import (
	"AoC-2024/utils"
	"slices"
	"strconv"
	"strings"
)

func Day05_2() {
	// input := strings.Split
	input := strings.Split(utils.GetData(5, utils.Real), "\n")

	// Seperate data
	page_order_rules := []string{}
	pages_to_produce := []string{}
	current_ref := &page_order_rules
	for _, line := range input {
		if line == "" {
			current_ref = &pages_to_produce
			continue
		}

		*current_ref = append(*current_ref, line)
	}

	pair_rules := make(map[int][]int)
	// Create rule pairs
	for _, pair := range page_order_rules {
		nums := strings.Split(pair, "|")
		key, _ := strconv.Atoi(nums[0])
		value, _ := strconv.Atoi(nums[1])
		pair_rules[key] = append(pair_rules[key], value)
	}

	// Convert updates from string to ints
	updates := [][]int{}
	for _, row := range pages_to_produce {
		spl_upd := strings.Split(row, ",")
		spl_upd_int := []int{}
		for _, c := range spl_upd {
			c_int, _ := strconv.Atoi(c)
			spl_upd_int = append(spl_upd_int, c_int)
		}
		updates = append(updates, spl_upd_int)
	}

	result := 0
	fixed_seq := [][]int{}
	for _, spl_upd := range updates {
		fixed := false
		for i := 0; i < len(spl_upd)-1; i++ {
			cur := &spl_upd[i]
			next := &spl_upd[i+1]

			if !slices.Contains(pair_rules[*cur], *next) {
				fixed = true
				if slices.Contains(pair_rules[*next], *cur) {
					temp := *cur
					*cur = *next
					*next = temp
					i = -1
				}
			}
		}
		if fixed {
			fixed_seq = append(fixed_seq, spl_upd)
		}
	}

	for _, i := range fixed_seq {
		mid_idx := len(i) / 2
		mid_val := i[mid_idx]
		result += mid_val
	}

	println(result)
}
