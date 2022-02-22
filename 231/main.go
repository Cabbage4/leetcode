package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(16))
}

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	for i := 1; i <= 32; i++ {
		if pow(i) == int64(n) {
			return true
		}
	}
	return false
}

func pow(n int) int64 {
	if n == 1 {
		return 2
	}

	if n%2 == 0 {
		t := pow(n / 2)
		return t * t
	}

	t := pow(n / 2)
	return t * t * 2
}
