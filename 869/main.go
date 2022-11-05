package main

import (
	"fmt"
)

func main() {
	fmt.Println(reorderedPowerOf2(46))
	fmt.Println(reorderedPowerOf2(47))
}

var mp = map[[10]int]bool{}

func init() {
	for n := 1; n <= 1e9; n = n << 1 {
		mp[count(n)] = true
	}
}

func reorderedPowerOf2(n int) bool {
	return mp[count(n)]
}

func count(n int) [10]int {
	var r [10]int
	for n != 0 {
		r[n%10]++
		n = n / 10
	}
	return r
}
