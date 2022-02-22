package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	if k > len(nums)-1 {
		k = len(nums) - 1
	}

	mp := make(map[int]int)
	for i := 0; i <= k; i++ {
		if mp[nums[i]] > 0 {
			return true
		}
		mp[nums[i]]++
	}

	for i := 1; i < len(nums)-k; i++ {
		mp[nums[i-1]]--
		if mp[nums[i-1]] == 0 {
			delete(mp, nums[i-1])
		}

		mp[nums[i+k]]++
		if mp[nums[i+k]] > 1 {
			return true
		}
	}

	return false
}
