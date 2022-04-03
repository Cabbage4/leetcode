package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFormArray([]int{15, 88}, [][]int{{88}, {15}}))
}

func canFormArray(arr []int, pieces [][]int) bool {
	mp := make(map[int][]int)
	for _, p := range pieces {
		mp[p[0]] = p
	}

	for i := 0; i < len(arr); {
		if len(mp) == 0 {
			return false
		}

		p, has := mp[arr[i]]
		if !has {
			return false
		}

		for k := 0; k < len(p); k++ {
			if i+k >= len(arr) || arr[i+k] != p[k] {
				return false
			}
		}

		i += len(p)
		delete(mp, p[0])
	}

	return len(mp) == 0
}
