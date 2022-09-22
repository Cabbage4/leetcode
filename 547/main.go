package main

func main() {
}

func findCircleNum(mp [][]int) int {
	u := make([]int, len(mp))
	for i := 0; i < len(mp); i++ {
		u[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if u[x] != x {
			u[x] = find(u[x])
		}
		return u[x]
	}

	union := func(x, y int) {
		u[find(x)] = find(y)
	}

	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {
			if mp[i][j] == 1 {
				union(i, j)
			}
		}
	}

	var r int
	for i := 0; i < len(u); i++ {
		if u[i] == i {
			r++
		}
	}

	return r
}
