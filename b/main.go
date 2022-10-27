package main

import "fmt"

func main() {
	//fmt.Println(subarrayGCD([]int{3, 3, 4, 1, 2}, 1))
	fmt.Println(subarrayGCD([]int{3, 12, 9, 6}, 3))
	fmt.Println(subarrayGCD([]int{4}, 7))
	fmt.Println(subarrayGCD([]int{9, 3, 1, 2, 6, 3}, 3))
}

func subarrayGCD(nums []int, k int) int {
	var r int
	for i := 0; i < len(nums); i++ {
		g := nums[i]
		for j := i; j < len(nums); j++ {
			g = gcd(g, nums[j])
			if g < k {
				break
			}
			if g == k {
				r++
			}
		}
	}
	return r
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
