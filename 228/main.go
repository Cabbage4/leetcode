package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{1, 3}))
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	res := make([]string, 0)

	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}

	a := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			if i == len(nums)-1 {
				res = append(res, fmt.Sprintf("%d->%d", nums[a], nums[i]))
				break
			}

			continue
		}

		if i-1 == a {
			res = append(res, fmt.Sprintf("%d", nums[a]))
			a = i

			if i == len(nums)-1 {
				res = append(res, fmt.Sprintf("%d", nums[i]))
				break
			}

			continue
		}

		res = append(res, fmt.Sprintf("%d->%d", nums[a], nums[i-1]))
		a = i

		if i == len(nums)-1 {
			res = append(res, fmt.Sprintf("%d", nums[i]))
			break
		}
	}

	return res
}
