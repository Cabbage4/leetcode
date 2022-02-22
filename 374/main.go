package main

func main() {

}
func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	left := 1
	right := n
	for {
		mid := (left + right) / 2
		ans := guess(mid)
		if ans == 0 {
			return mid
		} else if ans == -1 {
			right = mid
		} else {
			left = mid + 1
		}
	}
}
