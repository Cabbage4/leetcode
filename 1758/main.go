package main

import "fmt"

func main() {
	fmt.Println(minOperations("0100"))
	fmt.Println(minOperations("1111"))
}

func minOperations(s string) int {
	c1 := 0
	c2 := 0
	var l1 byte = '0'
	var l2 byte = '1'
	for i := 0; i < len(s); i++ {
		if s[i] != l1 {
			l1 = s[i]
		} else {
			l1 = (s[i]-'0'+1)%2 + '0'
			c1++
		}

		if s[i] != l2 {
			l2 = s[i]
		} else {
			l2 = (s[i]-'0'+1)%2 + '0'
			c2++
		}
	}

	if c1 < c2 {
		return c1
	}

	return c2
}
