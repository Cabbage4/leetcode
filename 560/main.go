package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}

func subarraySum(nums []int, k int) int {
	var r int
	var sum int
	mp := make(map[int]int)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := mp[sum-k]; ok {
			r += v
		}
		mp[sum]++
	}

	return r
}
