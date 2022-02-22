package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(16))
}

func isPowerOfFour(n int) bool {
	return ((n-1)&n) == 0 && ((n-1)%3 == 0)
}
