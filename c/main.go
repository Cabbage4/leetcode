package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strconv.FormatInt(25, 2))
	fmt.Println(strconv.FormatInt(72, 2))
	fmt.Println(minimizeXor(25, 72))
}

func minimizeXor(num1 int, num2 int) int {
	num2Count := count(num2)

	num1Byte := []byte(strconv.FormatInt(int64(num1), 2))
	if len(num1Byte) != 32 {
		num1Byte = append([]byte(strings.Repeat("0", 32-len(num1Byte))), num1Byte...)
	}
	targetByte := make([]byte, 32)
	for i := 0; i < 32; i++ {
		targetByte[i] = '0'
	}

	for i := 0; i < len(num1Byte); i++ {
		if num2Count == 0 {
			break
		}

		if num1Byte[i] == '1' {
			targetByte[i] = '1'
			num2Count--
		}
	}

	for i := len(targetByte) - 1; i >= 0; i-- {
		if num2Count == 0 {
			break
		}

		if targetByte[i] == '1' {
			continue
		}

		targetByte[i] = '1'
		num2Count--
	}

	r, _ := strconv.ParseInt(string(targetByte), 2, 64)
	return int(r)

}

func count(v int) int {
	r := 0
	for v > 0 {
		r += v & 1
		v = v >> 1
	}
	return r
}
