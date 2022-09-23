package main

func main() {
}

func arrayNesting(nums []int) int {
	var r int
	for i := range nums {
		c := 0
		for nums[i] != len(nums) {
			i, nums[i] = nums[i], len(nums)
			c++
		}

		if c > r {
			r = c
		}
	}
	return r
}
