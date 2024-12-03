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
}

func GetData(day int, dataType DataType) string {
	if dataType {
		return test_data[day]
	}
	requestURL := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, _ := http.NewRequest(http.MethodGet, requestURL, nil)

	dat, _ := os.ReadFile(".env")
	req.Header.Set("Cookie", string(dat))

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)

	return string(body)[:len(body)-1]
}
