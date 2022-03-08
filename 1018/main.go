package main

import "fmt"

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
}

func prefixesDivBy5(l []int) []bool {
	r := make([]bool, len(l))

	k := 0
	for i := 0; i < len(l); i++ {
		k = (k<<1 | l[i]) % 5
		r[i] = k == 0
	}

	return r
}
