package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 1}))
}

func singleNumber(nums []int) int {
	one := 0
	two := 0
	for _, v := range nums {
		one = (one ^ v) & ^two
		two = (two ^ v) & ^one
	}
	return one
}
