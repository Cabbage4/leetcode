package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPerfectNumber(1))
	fmt.Println(checkPerfectNumber(28))
	fmt.Println(checkPerfectNumber(7))
}

func checkPerfectNumber(num int) bool {
	list := []int{2, 3, 5, 7, 13, 17, 19, 31}
	for _, v := range list {
		if (2<<(v-1))*((2<<v)-1) == num {
			return true
		}
	}
	return false
}
