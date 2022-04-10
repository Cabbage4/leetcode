package main

import "fmt"

func main() {
	fmt.Println(mostFrequent([]int{1, 1, 2, 1, 2}, 1))
	fmt.Println(mostFrequent([]int{1, 100, 200, 1, 100}, 1))
}

func mostFrequent(nums []int, key int) int {
	mp := make([]int, 1001)
	for i, n := range nums {
		if n == key && i != len(nums)-1 {
			mp[nums[i+1]]++
		}
	}
	r := 0
	rc := 0
	for k, v := range mp {
		if v > rc {
			rc = v
			r = k
		}
	}
	return r
}
