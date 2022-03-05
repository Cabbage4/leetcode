package main

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
}

func reverseOnlyLetters(s string) string {
	sb := []byte(s)

	left := 0
	right := len(sb) - 1
	for left < right {
		for left < right && !is(sb[left]) {
			left++
		}
		if left >= right {
			break
		}

		for left < right && !is(sb[right]) {
			right--
		}
		if left >= right {
			break
		}

		sb[left], sb[right] = sb[right], sb[left]
		left++
		right--
	}

	return string(sb)
}

func is(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z'
}
