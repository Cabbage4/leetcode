package main

func main() {
}

func threeConsecutiveOdds(arr []int) bool {
	c := 0
	for i := 0; i < len(arr); i++ {
		if i < 2 {
			continue
		}

		if arr[i]%2 != 0 {
			c++
		} else {
			c = 0
		}
		if c >= 3 {
			return true
		}
	}

	return false
}
