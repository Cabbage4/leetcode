package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{1, 1, 1, 1, 2, 2, 2, 2}, 4))
	fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%k != 0 {
		return false
	}

	sort.Ints(nums)
	avg := sum / k
	if nums[len(nums)-1] > avg {
		return false
	}

	book := make([]bool, 1<<len(nums))
	var dfs func(int, int) bool
	dfs = func(state int, total int) bool {
		if state == 0 {
			return true
		}

		if book[state] {
			return false
		}
		book[state] = true

		for i := 0; i < len(nums); i++ {
			if state&(1<<i) == 0 {
				continue
			}

			if total+nums[i] > avg {
				break
			}

			if dfs(state&(^(1 << i)), (total+nums[i])%avg) {
				return true
			}
		}
		return false
	}

	return dfs(1<<len(nums)-1, 0)
}
