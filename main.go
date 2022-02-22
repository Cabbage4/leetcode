package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(readBinaryWatch(1))
}

func readBinaryWatch(turnedOn int) []string {
	book := make([]bool, 10)
	resultMp := make(map[string]bool)
	dfs(book, turnedOn, 0, resultMp)

	result := make([]string, 0)
	for k := range resultMp {
		if k == "" {
			continue
		}
		result = append(result, k)
	}

	return result
}

func dfs(book []bool, total int, cur int, result map[string]bool) {
	if cur == total {
		result[bookToStr(book)] = true
		return
	}

	for i := range book {
		if book[i] {
			continue
		}

		book[i] = true
		dfs(book, total, cur+1, result)
		book[i] = false
	}
}

func bookToStr(book []bool) string {
	result := bytes.NewBufferString("")

	top := 0
	for i := 0; i < 4; i++ {
		if book[i] {
			top += int(math.Pow(float64(2), float64(3-i)))
		}
	}

	if top >= 12 {
		return ""
	}

	result.WriteString(strconv.Itoa(top))
	result.WriteString(":")

	low := 0
	for i := 4; i < 10; i++ {
		if book[i] {
			low += int(math.Pow(float64(2), float64(i-4)))
		}
	}

	if low >= 60 {
		return ""
	}

	if low < 10 {
		result.WriteString("0")
	}
	result.WriteString(strconv.Itoa(low))

	return result.String()
}
