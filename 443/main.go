package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	target := []byte("aabbccc")
	fmt.Println(compress(target), string(target))
}

func compress(chars []byte) int {
	var r int
	for i := 0; i < len(chars); {
		j := i + 1
		for j < len(chars) && chars[i] == chars[j] {
			j++
		}

		chars[r] = chars[i]
		r++

		if j-i > 1 {
			for _, v := range strings.Split(strconv.Itoa(j-i), "") {
				chars[r] = v[0]
				r++
			}
		}

		i = j
	}

	return r
}
