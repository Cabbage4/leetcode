package main

import "fmt"

func main() {
	testList := []int{3, 2, 2, 3}
	fmt.Println(removeElement(testList, 3), testList)
	testList = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(testList, 2), testList)
	testList = []int{4, 5}
	fmt.Println(removeElement(testList, 5), testList)
}

func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}

	return right
}
