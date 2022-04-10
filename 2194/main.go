package main

import "fmt"

func main() {
	fmt.Println(cellsInRange("K1:L2"))
}

func cellsInRange(s string) []string {
	r1, r2 := s[0], s[3]
	c1, c2 := s[1], s[4]

	r := make([]string, 0)
	for i := r1; i <= r2; i++ {
		for j := c1; j <= c2; j++ {
			r = append(r, string(i)+string(j))
		}
	}
	return r
}
