package main

import "fmt"

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	if target == 'z' {
		return letters[0]
	}

	left := 0
	right := len(letters)

	for left < right {
		mid := (left + right) / 2
		if letters[mid] == target || letters[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left >= len(letters) {
		return letters[0]
	}

	return letters[left]
}
