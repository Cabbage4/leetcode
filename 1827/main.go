package main

func main() {
}

func minOperations(nums []int) int {
	r := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}

		r += nums[i-1] - nums[i] + 1
		nums[i] = nums[i-1] + 1
	}
	return r
}
