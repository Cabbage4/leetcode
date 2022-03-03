package main

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
}

func largeGroupPositions(s string) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(s); {
		j := i + 1
		for j < len(s) && s[i] == s[j] {
			j++
		}

		if j-i >= 3 {
			result = append(result, []int{i, j - 1})
		}
		i = j
	}
	return result
}
