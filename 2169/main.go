package main

import "fmt"

func main() {
	fmt.Println(countOperations(2, 3))
}

func countOperations(num1 int, num2 int) int {
	if num2 == 0 || num1 == 0 {
		return 0
	}

	return num1/num2 + countOperations(num2, num1%num2)
}
