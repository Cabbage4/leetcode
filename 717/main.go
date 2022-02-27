package main

import "fmt"

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
}

func isOneBitCharacter(bits []int) bool {
	if bits[len(bits)-1] == 1 {
		return false
	}

	if len(bits) < 2 {
		return true
	}

	bits = bits[:len(bits)-1]

	for i := 0; i < len(bits); {
		if bits[i] == 0 {
			i++
		} else {
			i = i + 2
			if i > len(bits) {
				return false
			}
		}
	}

	return true
}
