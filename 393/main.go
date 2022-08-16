package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(validUtf8([]int{10}) == true)
	//fmt.Println(validUtf8([]int{145}) == false)
	//fmt.Println(validUtf8([]int{197, 130, 1}) == true)
	//fmt.Println(validUtf8([]int{235, 140, 4}) == false)
	//fmt.Println(validUtf8([]int{230, 136, 145}) == true)
	fmt.Println(validUtf8([]int{250, 145, 145, 145, 145}) == false)
}

func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		if data[i] < 128 {
			i++
			continue
		}

		firstDataStr := strconv.FormatInt(int64(data[i]), 2)

		j := 1
		for ; j < len(firstDataStr); j++ {
			if firstDataStr[j] == '0' {
				break
			}

			if i+j >= len(data) {
				return false
			}

			if data[i+j] < 128 || data[i+j] >= 192 {
				return false
			}
		}

		if j == 1 || j > 4 {
			return false
		}

		i += j
	}

	return true
}
