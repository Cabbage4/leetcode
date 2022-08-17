package main

import "fmt"

func main() {
	//fmt.Println(integerReplacement(8))
	fmt.Println(integerReplacement(7))
	//fmt.Println(integerReplacement(4))
}

func integerReplacement(n int) int {
	return dfs(n, make(map[int]int))
}

func dfs(n int, mp map[int]int) int {
	if n == 1 {
		return 0
	}

	if r, ok := mp[n]; ok {
		return r
	}

	if n%2 == 0 {
		r := dfs(n/2, mp) + 1
		mp[n] = r
		return r
	}

	a := dfs(n+1, mp) + 1
	b := dfs(n-1, mp) + 1

	if a > b {
		mp[n] = b
		return b
	}

	mp[n] = a
	return a
}
