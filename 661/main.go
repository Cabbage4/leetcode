package main

func main() {

}

var direct = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{1, -1},
	{-1, -1},
	{-1, 1},
}

func imageSmoother(img [][]int) [][]int {
	row := len(img)
	col := len(img[0])
	result := make([][]int, row)

	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
		for j := 0; j < col; j++ {
			cur := 1
			total := img[i][j]
			for _, v := range direct {
				x := i + v[0]
				y := j + v[1]
				if x >= 0 && x < row && y >= 0 && y < col {
					cur++
					total += img[x][y]
				}
			}

			result[i][j] = total / cur
		}
	}

	return result
}
