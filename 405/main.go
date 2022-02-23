package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(toHex(26))
}

func toHex(num int) string {
	if num < 0 {
		num = int(math.Pow(2, 32)) + num
	}

	mp := make(map[int]string)
	for i := 0; i <= 9; i++ {
		mp[i] = string('0' + i)
	}

	for i := 10; i <= 15; i++ {
		mp[i] = string('a' + i - 10)
	}

	list := make([]string, 8)
	for i := 0; i < 8; i++ {
		list[7-i] = mp[num&15]
		num = num >> 4
	}

	for i := 7; i >= 0; i-- {
		if list[0] != "0" {
			break
		}

		if len(list) == 1 {
			break
		}

		list = list[1:]
	}

	return strings.Join(list, "")
}
