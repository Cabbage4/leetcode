package main

func main() {

}

func projectionArea(grid [][]int) int {
	n := len(grid)
	r := 0

	for i := 0; i < n; i++ {
		m1 := 0
		m2 := 0
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				r++
			}
			if grid[i][j] > m1 {
				m1 = grid[i][j]
			}
			if grid[j][i] > m2 {
				m2 = grid[j][i]
			}
		}
		r += m1 + m2
	}

	return r
}
