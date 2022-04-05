package main

func main() {
}

func findCenter(edges [][]int) int {
	n11, n12 := edges[0][0], edges[0][1]
	n21, n22 := edges[1][0], edges[1][1]
	if n11 == n21 || n11 == n22 {
		return n11
	}
	if n12 == n21 || n12 == n22 {
		return n12
	}

	return -1
}
