package main

import (
	"fmt"
)

func main() {
	fmt.Println(getLucky("dbvmfhnttvr", 5))
	fmt.Println(getLucky("iiii", 1))
	fmt.Println(getLucky("leetcode", 2))
}

func getLucky(s string, k int) int {
	pr := 0
	for i := 0; i < len(s); i++ {
		t := int(s[i] - 'a' + 1)
		pr += t/10 + t%10
	}

	r := pr
	tr := 0
	for i := 0; i < k-1; i++ {
		for pr > 0 {
			tr += pr % 10
			pr = pr / 10
		}

		if i == k-1-1 {
			r = tr
			break
		}

		pr = tr
		tr = 0
	}

	return r
}
