package main

import "math"

func main() {

}

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return 1
	}

	ln := int(math.Pow10(4) + 1)

	mp := make([]int, ln)
	for _, v := range trust {
		mp[v[1]]++
		mp[v[0]]--
	}

	for k, v := range mp {
		if v == n-1 {
			return k
		}
	}
	return -1
}
