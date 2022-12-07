package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("())"))
	fmt.Println(longestValidParentheses("()((())"))
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("(()"))
}

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		if s[stack[len(stack)-1]] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}

	if len(stack) == 0 {
		return len(s)
	}

	r := stack[0]
	for i := 1; i < len(stack); i++ {
		tr := stack[i] - 1 - stack[i-1]
		if tr > r {
			r = tr
		}
	}

	if stack[len(stack)-1] != len(s)-1 {
		tr := len(s) - 1 - stack[len(stack)-1]
		if tr > r {
			r = tr
		}
	}

	return r
}
