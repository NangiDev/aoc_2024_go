package day05

import (
	"AoC-2024/utils"
	"slices"
	"strconv"
	"strings"
)

func Day05_1() {
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
			c_int, _ := strconv.Atoi(string(c))
			spl_upd_int = append(spl_upd_int, c_int)
		}
		updates = append(updates, spl_upd_int)
	}

	result := 0
	for _, spl_upd := range updates {
		valid := true
		for i := 0; i < len(spl_upd)-1; i++ {
			cur := spl_upd[i]
			next := spl_upd[i+1]

			if !slices.Contains(pair_rules[cur], next) {
				valid = false
				break
			}
		}
		if valid {
			mid_idx := len(spl_upd) / 2
			mid_val := spl_upd[mid_idx]
			result += mid_val
		}
	}

	println(result)
}
