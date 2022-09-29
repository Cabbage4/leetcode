package main

import (
	"fmt"
	"sort"
)

func main() {
	data := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
	//data := [][]string{{"David", "David0@m.co", "David1@m.co"}, {"David", "David3@m.co", "David4@m.co"}, {"David", "David4@m.co", "David5@m.co"}, {"David", "David2@m.co", "David3@m.co"}, {"David", "David1@m.co", "David2@m.co"}}
	fmt.Println(accountsMerge(data))
}

func accountsMerge(accounts [][]string) [][]string {
	parent := make([]int, len(accounts))
	for i := 0; i < len(accounts); i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(i, j int) {
		x, y := find(i), find(j)
		if x == y {
			return
		}

		parent[x] = y
	}

	mp := make(map[string][]int)
	for i := 0; i < len(accounts); i++ {
		for _, a := range accounts[i][1:] {
			mp[a] = append(mp[a], i)
		}
	}

	for _, v := range mp {
		for i := 1; i < len(v); i++ {
			union(v[0], v[i])
		}
	}

	mp2 := make(map[int][]string)
	mp2H := make(map[int]map[string]bool)
	for i := 0; i < len(parent); i++ {
		x := find(i)

		for _, s := range accounts[i][1:] {
			if mp2H[x] == nil {
				mp2H[x] = make(map[string]bool)
			}
			if !mp2H[x][s] {
				mp2[x] = append(mp2[x], s)
				mp2H[x][s] = true
			}
		}
	}

	var r [][]string
	for index, list := range mp2 {
		sort.Strings(list)
		r = append(r, append([]string{accounts[index][0]}, list...))
	}

	return r
}
