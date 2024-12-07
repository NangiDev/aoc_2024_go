package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type DataType bool

const (
	Test DataType = true
	Real DataType = false
)

var test_data = map[int]string{
	1: `3   4
4   3
2   5
1   3
3   9
3   3`,
	2: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
	3: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	4: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
	5: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`,
	6: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
}

func GetData(day int, dataType DataType) string {
	if dataType {
		return test_data[day]
	}

	filepath := fmt.Sprintf("./day%02d/input.txt", day)
	if _, err := os.Stat(filepath); err == nil {
		dat, _ := os.ReadFile(filepath)
		return string(dat)[:len(dat)-1]
	}

	requestURL := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, _ := http.NewRequest(http.MethodGet, requestURL, nil)

	dat, _ := os.ReadFile(".env")
	req.Header.Set("Cookie", string(dat))

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	os.WriteFile(filepath, body, os.ModePerm)

	return string(body)[:len(body)-1]
}
