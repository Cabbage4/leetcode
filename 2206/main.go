package main

func main() {
}

func divideArray(nums []int) bool {
	mp := make([]int, 501)
	for _, v := range nums {
		mp[v]++
	}
	for _, v := range mp {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
