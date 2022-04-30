package main

func main() {
}

func singleNumber(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	mask := xor & (-xor)
	r := []int{0, 0}
	for _, v := range nums {
		if v&mask == 0 {
			r[0] ^= v
		} else {
			r[1] ^= v
		}
	}
	return r
}
