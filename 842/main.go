package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(splitIntoFibonacci("0000"))
	fmt.Println(splitIntoFibonacci("74912134825162255812723932620170946950766784234934"))
	fmt.Println(splitIntoFibonacci("1101111"))
	fmt.Println(splitIntoFibonacci("020406"))
}

func splitIntoFibonacci(num string) []int {
	hasLeadZero := func(s string) bool {
		if len(s) == 1 {
			return false
		}

		if s[0] == '0' {
			return true
		}

		return false
	}

	var r []int
	var dfs func(int, int, string) bool
	dfs = func(i1, i2 int, s string) bool {
		if i3, _ := strconv.Atoi(s); !hasLeadZero(s) && i3 <= math.MaxInt32 && i1+i2 == i3 {
			r = append([]int{i3}, r...)
			return true
		}

		for i := 1; i < len(s); i++ {
			if hasLeadZero(s[:i]) {
				continue
			}

			i3, _ := strconv.Atoi(s[:i])
			if i3 > math.MaxInt32 {
				break
			}

			if i1+i2 != i3 {
				continue
			}

			if dfs(i2, i3, s[i:]) {
				r = append([]int{i3}, r...)
				return true
			}
		}

		return false
	}

	for i := 1; i < len(num)-2; i++ {
		if hasLeadZero(num[:i]) {
			break
		}

		i1, _ := strconv.Atoi(num[:i])
		if i1 > math.MaxInt32 {
			break
		}

		for j := i + 1; j < len(num)-1; j++ {
			if hasLeadZero(num[i:j]) {
				break
			}

			i2, _ := strconv.Atoi(num[i:j])
			if i2 > math.MaxInt32 {
				break
			}

			if dfs(i1, i2, num[j:]) {
				r = append([]int{i1, i2}, r...)
				return r
			}
		}
	}

	return r
}
