package main

func main() {

}

func surfaceArea(grid [][]int) int {
	n := len(grid)
	r := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				r = 4*grid[i][j] + 2 + r
			}

			if i > 0 {
				r = r - min(grid[i][j], grid[i-1][j])*2
			}

			if j > 0 {
				r = r - min(grid[i][j], grid[i][j-1])*2
			}
		}
	}

	return r
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
