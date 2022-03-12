package main

func main() {
}

func numIdenticalPairs(nums []int) int {
	mp := make([]int, 101)

	r := 0
	for _, v := range nums {
		mp[v]++
		if mp[v] > 1 {
			r += mp[v] - 1
		}
	}

	return r
}
