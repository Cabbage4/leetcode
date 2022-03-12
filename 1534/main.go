package main

func main() {
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	r := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i], arr[j]) > a {
				continue
			}

			for k := j + 1; k < len(arr); k++ {
				if abs(arr[j], arr[k]) > b {
					continue
				}

				if abs(arr[i], arr[k]) <= c {
					r++
				}
			}
		}
	}
	return r
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
