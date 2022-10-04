package main

import "fmt"

func main() {
	fmt.Println(monotoneIncreasingDigits(10))
	fmt.Println(monotoneIncreasingDigits(1234))
	fmt.Println(monotoneIncreasingDigits(332))
}

func monotoneIncreasingDigits(n int) int {
	list := make([]int, 0)
	tn := n
	for tn > 0 {
		list = append(list, tn%10)
		tn /= 10
	}

	index := -1
	for i := len(list) - 1; i > 0; i-- {
		if list[i] <= list[i-1] {
			continue
		}
		index = i
		break
	}

	if index == -1 {
		return n
	}

	for index < len(list)-1 && list[index]-1 < list[index+1] {
		index++
	}

	for i := index - 1; i >= 0; i-- {
		list[i] = 9
	}
	list[index]--

	r := 0
	for i := len(list) - 1; i >= 0; i-- {
		r = r*10 + list[i]
	}

	return r
}
