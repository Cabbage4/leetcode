package main

import "fmt"

func main() {
	fmt.Println(numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
}

func numberOfBoomerangs(points [][]int) int {

	mp := make([]map[int]int, len(points))
	for i := range mp {
		mp[i] = make(map[int]int)
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dis := distance(points[i], points[j])
			mp[i][dis]++
			mp[j][dis]++
		}
	}

	var r int
	for _, mmp := range mp {
		for _, v := range mmp {
			if v >= 2 {
				r += v * (v - 1)
			}
		}
	}

	return r
}

func distance(a []int, b []int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return x*x + y*y
}
