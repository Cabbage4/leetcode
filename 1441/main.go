package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
}

func buildArray(target []int, n int) []string {
	r := make([]string, 0)

	ti := 0
	for i := 1; i <= n; i++ {
		if target[ti] > i {
			r = append(r, "Push", "Pop")
			continue
		}

		r = append(r, "Push")
		ti++
		if ti >= len(target) {
			break
		}
	}

	return r
}
