package main

import "fmt"

func main() {
	n := 15
	r := generateMatrix(n)
	for i := 0; i < n; i++ {
		fmt.Println(r[i])
	}
}

func generateMatrix(n int) [][]int {
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
	}

	c := 1
	for i := 0; i <= n/2; i++ {
		for k := i; k < n-i; k++ {
			r[i][k] = c
			c++
		}

		for k := i + 1; k < n-i; k++ {
			r[k][n-i-1] = c
			c++
		}

		for k := n - i - 2; k >= i; k-- {
			r[n-i-1][k] = c
			c++
		}

		for k := n - i - 2; k >= i+1; k-- {
			r[k][i] = c
			c++
		}
	}

	return r
}
