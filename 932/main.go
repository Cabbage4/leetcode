package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(beautifulArray(i))
	}
}

func beautifulArray(n int) []int {
	if n == 1 {
		return []int{1}
	}

	r := make([]int, n)
	if n%2 == 0 {
		list := beautifulArray(n / 2)
		mid := n / 2
		for i := 0; i < mid; i++ {
			r[i] = 2*list[i] - 1
			r[i+mid] = 2 * list[i]
		}
	} else {
		leftList := beautifulArray((n + 1) / 2)
		rightList := beautifulArray((n - 1) / 2)

		for i := 0; i < len(leftList); i++ {
			r[i] = 2*leftList[i] - 1
		}

		for i := 0; i < len(rightList); i++ {
			r[i+len(leftList)] = 2 * rightList[i]
		}
	}

	return r
}
