package main

func main() {

}

func search(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}
