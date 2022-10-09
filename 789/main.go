package main

func main() {
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	x := abs(target[0]) + abs(target[1])
	for _, g := range ghosts {
		y := abs(g[0]-target[0]) + abs(g[1]-target[1])
		if y <= x {
			return false
		}
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
