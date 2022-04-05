package main

import "fmt"

func main() {
	fmt.Println(countTriples(5))
}

func countTriples(n int) int {
	set := make([]bool, n*n+1)
	for i := 2; i <= n; i++ {
		set[i*i] = true
	}

	r := 0
	for i := 1; i <= n; i++ {
		i2 := i * i
		for j := i; i2+j*j <= n*n; j++ {
			j2 := j * j
			if set[i2+j2] {
				r++
			}
		}
	}
	return 2 * r
}
