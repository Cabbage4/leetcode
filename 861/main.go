package main

func main() {
}

func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	r := (1 << (n - 1)) * m
	for j := 1; j < n; j++ {
		count := 0
		for _, v := range grid {
			if v[j] == v[0] {
				count++
			}
		}
		if m-count > count {
			count = m - count
		}
		r += (1 << (n - 1 - j)) * count
	}
	return r
}
