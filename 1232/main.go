package main

func main() {
}

func checkStraightLine(c [][]int) bool {
	if len(c) < 3 {
		return true
	}

	x1, x2 := c[0][0], c[1][0]
	y1, y2 := c[0][1], c[1][1]

	for i := 2; i < len(c); i++ {
		if (y1-y2)*(x1-c[i][0]) != (x1-x2)*(y1-c[i][1]) {
			return false
		}
	}

	return true
}
