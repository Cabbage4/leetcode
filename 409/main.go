package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}

func longestPalindrome(s string) int {
	mp := make([]int, 256)
	for _, v := range s {
		mp[v]++
	}

	result := 0
	maxIndex := -1
	for k, v := range mp {
		if v == 0 {
			continue
		}

		if v%2 == 0 {
			result += v
			mp[k] = 0
		} else {
			if maxIndex == -1 {
				maxIndex = k
			} else {
				if mp[maxIndex] < v {
					maxIndex = k
				}
			}
		}
	}

	if maxIndex != -1 {
		result += mp[maxIndex]
		mp[maxIndex] = 0
	}

	for _, v := range mp {
		if v != 0 {
			result += v - 1
		}
	}

	return result
}
