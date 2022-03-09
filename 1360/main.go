package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
	fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31"))
}

func daysBetweenDates(date1 string, date2 string) int {
	return abs(tran(date1), tran(date2))
}

var mp = []int{
	0, 0,
	31, 28, 31,
	30, 31, 30,
	31, 31, 30,
	31, 30, 0,
}

var yp = func() map[int]bool {
	r := make(map[int]bool)
	for y := 1971; y <= 2100; y++ {
		if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
			r[y] = true
		}
	}
	return r
}()

func tran(date string) int {
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

		v, _ := strconv.Atoi(date[i:j])
		g = append(g, v)
		i = j
	}

	y, m, d := g[0], g[1], g[2]
	r := d
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m >= 3 {
		r++
	}
	for i := 1; i <= m; i++ {
		r += mp[i]
	}

	for i := 1972; i <= y-1; i++ {
		if yp[i] {
			r++
		}
	}

	return r + (y-1)*365
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
