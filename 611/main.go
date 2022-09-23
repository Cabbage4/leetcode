package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(triangleNumber([]int{4, 2, 3, 4}))
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)

	binarySearch := func(left, right int, target int, flag bool) int {
		for left < right {
			mid := (left + right) / 2
			if nums[mid] < target {
				left = mid + 1
			}

			if nums[mid] > target {
				right = mid
			}

			if nums[mid] == target {
				if flag {
					right = mid
				} else {
					left = mid + 1
				}
			}
		}

		return left
	}

	var r int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			min := nums[j] - nums[i] + 1
			max := nums[j] + nums[i] - 1
			minIndex := binarySearch(j+1, len(nums), min, true)
			maxIndex := binarySearch(j+1, len(nums), max, false)

			if maxIndex > minIndex {
				r += maxIndex - minIndex
			}
		}
	}

	return r
}
