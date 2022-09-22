package main

func main() {
}

func singleNonDuplicate(nums []int) int {
	var r int
	for _, v := range nums {
		r ^= v
	}
	return r
}
