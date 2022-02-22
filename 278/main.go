package main

func main() {
}

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	mid := (1 + n) / 2
	for left < right {
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return left
}
