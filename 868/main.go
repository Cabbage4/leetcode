package main

import "fmt"

func main() {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(8))
	fmt.Println(binaryGap(5))
}

func binaryGap(n int) int {
	result := 0
	last := -1
	index := 0

	for n > 0 {
		v := n & 1

		if v == 1 {
			if last != -1 {
				result = max(result, index-last)
			}

			last = index
		}

		index++
		n = n >> 1
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
