package main

import "fmt"

func main() {
	fmt.Println(countEven(4))
}

func countEven(target int) int {
	c := 0
	num := target
	for num > 0 {
		c += num % 10
		num = num / 10
	}

	if c%2 == 0 {
		return target / 2
	}

	if target%10 == 0 {
		return target/2 - 1
	}
	return (target - 1) / 2
}
