package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	var r []int
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		list := make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&list[j])
		}

		var c int
		fmt.Scan(&c)

		index := 0
		var indexValue int64 = 0
		for j := 0; j < c; j++ {
			var a, b int64
			fmt.Scanf("%d %d", &a, &b)

			if a == 1 {
				list = append(list, b)
				continue
			}

			indexValue += b
			if list[index] > indexValue {
				r = append(r, index+1)
				continue
			}

			for list[index] <= indexValue {
				indexValue = indexValue - list[index]
				index = (index + 1 + len(list)) % len(list)
			}

			r = append(r, index+1)
		}
	}

	for _, v := range r {
		fmt.Println(v)
	}
}
