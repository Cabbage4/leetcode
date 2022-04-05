package main

func main() {
}

func arraySign(nums []int) int {
	r := true

	for _, v := range nums {
		if v == 0 {
			return 0
		}

		if v < 0 {
			r = !r
		}
	}

	if r {
		return 1
	}
	return -1
}
