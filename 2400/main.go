package main

import "fmt"

func main() {
	fmt.Println(numberOfWays(1, 2, 3))
}

func numberOfWays(startPos int, endPos int, k int) int {
	return h(startPos, endPos, k, make(map[string]int))
}

func h(s, e, k int, mp map[string]int) int {
	if v, ok := mp[fmt.Sprintf("%d_%d_%d", s, e, k)]; ok {
		return v
	}

	if k < abs(e, s) || (k-abs(e, s))%2 != 0 {
		mp[fmt.Sprintf("%d_%d_%d", s, e, k)] = 0
		return 0
	}

	if s == e && k == 0 {
		return 1
	}

	r := (h(s+1, e, k-1, mp) + h(s-1, e, k-1, mp)) % 1000000007
	mp[fmt.Sprintf("%d_%d_%d", s, e, k)] = r
	return r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
