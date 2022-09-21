package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1}))
	fmt.Println(findSubsequences([]int{4, 4, 3, 2, 1}))
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}

func findSubsequences(nums []int) [][]int {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = -1
	for i := len(nums) - 2; i >= 0; i-- {
		j := i + 1
		for j < len(nums) && nums[j] <= nums[i] && dp[j] != -1 {
			j = dp[j]
		}

		if nums[j] >= nums[i] {
			dp[i] = j
			continue
		}

		dp[i] = -1
	}

	var r [][]int

	set := make(map[string]bool)
	var dfs func([]int, int, []int)
	dfs = func(list []int, curIndex int, curList []int) {
		if curIndex == len(list) {
			if len(curList) != 1 {
				key := ""
				for _, v := range curList {
					key += "_" + strconv.Itoa(v)
				}
				if !set[key] {
					r = append(r, curList)
					set[key] = true
				}
			}
			return
		}

		for i := curIndex; i < len(list); i++ {
			tmpList2 := make([]int, len(curList)+1)
			tmpList2[len(tmpList2)-1] = list[i]
			copy(tmpList2, curList)
			dfs(list, i+1, tmpList2)

			tmpList1 := make([]int, len(curList))
			copy(tmpList1, curList)
			dfs(list, i+1, tmpList1)
		}
	}

	for i := 0; i < len(dp); i++ {
		j := dp[i]
		curList := []int{nums[i]}
		for j != -1 {
			curList = append(curList, nums[j])
			j = dp[j]
		}

		dfs(curList, 1, []int{nums[i]})
	}

	return r
}
