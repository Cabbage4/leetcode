package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findComplement(5))
}

func findComplement(num int) int {
	vb := bytes.NewBufferString("")
	b := strconv.FormatInt(int64(num), 2)
	for _, v := range b {
		vb.WriteString(string('0' + '1' - v))
	}
	result, _ := strconv.ParseInt(vb.String(), 2, 32)
	return int(result)
}
