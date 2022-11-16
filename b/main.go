package main

import "fmt"

func main() {
	fmt.Println(maximumSubarraySum([]int{9, 9, 9, 1, 2, 3}, 3))
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
	fmt.Println(maximumSubarraySum([]int{4, 4, 4}, 3))
}

func maximumSubarraySum(nums []int, k int) int64 {
	var r int64

	var sum int64 = 0
	dupMap := make(map[int]bool)
	cmp := make(map[int]int)
	for i := 0; i < k; i++ {
		sum += int64(nums[i])
		cmp[nums[i]]++

		if cmp[nums[i]] >= 2 {
			dupMap[nums[i]] = true
		}
	}

	if len(dupMap) == 0 {
		r = sum
	}

	for i := k; i < len(nums); i++ {
		sum = sum + int64(nums[i]) - int64(nums[i-k])

		cmp[nums[i-k]]--
		cmp[nums[i]]++

		if cmp[nums[i-k]] < 2 {
			delete(dupMap, nums[i-k])
		}

		if cmp[nums[i]] >= 2 {
			dupMap[nums[i]] = true
		}

		if len(dupMap) == 0 && sum > r {
			r = sum
		}
	}

	return r
}
