package main

import "fmt"

func main() {
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0}, 3))
	fmt.Println(numSubarraysWithSum([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1))
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 1, 0, 0}, 0))
}

func numSubarraysWithSum(nums []int, goal int) int {
	var r int
	sum := 0
	mp := make(map[int]int)
	for i := range nums {
		mp[sum]++
		sum += nums[i]
		r += mp[sum-goal]
	}
	return r
}
