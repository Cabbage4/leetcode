package main

func main() {
}

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr)

	for left < right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
