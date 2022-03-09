package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
}

func dayOfYear(date string) int {
	cmp := []int{
		0, 0,
		31, 28, 31,
		30, 31, 30,
		31, 31, 30,
		31, 30,
	}

	g := make([]int, 0)
	for i := 0; i < len(date); {
		if date[i] == '-' {
			i++
			continue
		}
		j := i + 1
		for j < len(date) {
			if date[j] != '-' {
				j++
				continue
			}
			break
		}

		t, _ := strconv.Atoi(date[i:j])
		g = append(g, t)
		i = j
	}

	y, m, d := g[0], g[1], g[2]
	r := d
	if ((y%100 != 0 && y%4 == 0) || y%400 == 0) && m > 2 {
		r++
	}
	for i := 1; i <= m; i++ {
		r += cmp[i]
	}
	return r
}
