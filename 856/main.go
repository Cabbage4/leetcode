package main

import "fmt"

func main() {
	fmt.Println(scoreOfParentheses("()"))
	fmt.Println(scoreOfParentheses("(()())"))
	fmt.Println(scoreOfParentheses("()()"))
}

func scoreOfParentheses(s string) int {
	stack := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' {
			stack = append(stack, -1)
			continue
		}

		if stack[len(stack)-1] == -1 {
			stack[len(stack)-1] = 1
			continue
		}

		count := 0
		for len(stack) != 0 && stack[len(stack)-1] != -1 {
			count += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack[len(stack)-1] = 2 * count
	}

	var r int
	for i := range stack {
		r += stack[i]
	}

	return r
}
