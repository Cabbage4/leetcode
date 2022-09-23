package main

func main() {
}

func leastBricks(wall [][]int) int {
	mp := make(map[int]int)
	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j < len(wall[i])-1; j++ {
			sum += wall[i][j]
			mp[sum]++
		}
	}

	var r int
	for _, v := range mp {
		if v > r {
			r = v
		}
	}

	return len(wall) - r
}
