package main

func main() {
}

func countPairs(nums []int, k int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (j*i)%k != 0 {
				continue
			}

			if nums[i] == nums[j] {
				r++
			}
		}
	}
	return r
}
