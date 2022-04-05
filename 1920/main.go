package main

func main() {
}

func buildArray(nums []int) []int {
	r := make([]int, len(nums))
	for i := range nums {
		r[i] = nums[nums[i]]
	}
	return r
}
