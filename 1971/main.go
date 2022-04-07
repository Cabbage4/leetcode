package main

import "fmt"

func main() {
	//test := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	//fmt.Println(validPath(6, test, 0, 5))
	//fmt.Println(validPath(6, test, 1, 2))

	//test := [][]int{}
	//fmt.Println(validPath(1, test, 0, 0))

	test := [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}
	fmt.Println(validPath(10, test, 7, 5))
}

func validPath(n int, edges [][]int, s int, d int) bool {
	if s == d {
		return true
	}
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}

	for _, e := range edges {
		a := find(list, e[0])
		b := find(list, e[1])
		list[a] = b
	}

	return find(list, s) == find(list, d)
}

func find(list []int, target int) int {
	if list[target] == target {
		return target
	}

	list[target] = find(list, list[target])
	return list[target]
}
