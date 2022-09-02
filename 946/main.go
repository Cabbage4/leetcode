package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))

	pushedIndex := 0
	poppedIndex := 0
	for {
		for len(stack) != 0 && popped[poppedIndex] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			poppedIndex++
		}

		if poppedIndex == len(popped) {
			return len(stack) == 0 && pushedIndex == len(pushed)
		}

		if pushedIndex == len(pushed) {
			return len(stack) == 0 && poppedIndex == len(popped)
		}

		for pushedIndex < len(pushed) {
			stack = append(stack, pushed[pushedIndex])
			pushedIndex++

			if pushed[pushedIndex-1] == popped[poppedIndex] {
				break
			}
		}
	}
}
