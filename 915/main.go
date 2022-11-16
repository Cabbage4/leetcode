package main

func main() {
}

func partitionDisjoint(nums []int) int {
	leftPos, leftMax := 0, nums[0]
	curMax := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		curMax = max(curMax, nums[i])
		if nums[i] < leftMax {
			leftMax = curMax
			leftPos = i
		}
	}
	return leftPos + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
