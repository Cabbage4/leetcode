package main

import (
	"fmt"
)

func main() {
	fmt.Println(secondHighest("ck077"))
	fmt.Println(secondHighest("dfa12321afd"))
}

func secondHighest(s string) int {
	max1 := -1
	max2 := -1
	for i := range s {
		if !('0' <= s[i] && s[i] <= '9') {
			continue
		}

		t := int(s[i] - '0')
		if t == max1 || t == max2 {
			continue
		}

		if t > max1 {
			max2 = max1
			max1 = t
			continue
		}

		if t > max2 {
			max2 = t
		}
	}

	return max2
}
