package main

import "fmt"

func main() {
	fmt.Println(twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	mp := make([]int, 101)
	mp1 := make([]int, 101)
	mp2 := make([]int, 101)
	mp3 := make([]int, 101)

	for _, v := range nums1 {
		mp1[v] = 1
		mp[v]++
	}
	for _, v := range nums2 {
		mp2[v] = 1
		mp[v]++
	}
	for _, v := range nums3 {
		mp3[v] = 1
		mp[v]++
	}

	r := make([]int, 0)
	for k := range mp {
		if mp1[k]+mp2[k]+mp3[k] >= 2 {
			r = append(r, k)
		}
	}

	return r
}
