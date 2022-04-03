package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(reformatNumber("--17-5 229 35-39475 "))
	fmt.Println(reformatNumber("123 4-567"))
	fmt.Println(reformatNumber("123 4-5678"))
}

func reformatNumber(number string) string {
	dl := bytes.NewBufferString("")
	for i := 0; i < len(number); i++ {
		if number[i] == ' ' || number[i] == '-' {
			continue
		}

		dl.WriteByte(number[i])
	}

	r := bytes.NewBufferString("")

	i := 0
	ds := dl.String()
	for i < len(ds)-3 {
		r.WriteString(ds[i : i+3])
		r.WriteByte('-')

		if i+3 >= len(ds) {
			break
		}

		i = i + 3
	}

	if len(ds)-i == 3 {
		r.WriteString(ds[i : i+3])
	} else if len(ds)-i == 2 {
		r.WriteString(ds[i : i+2])
	} else if len(ds)-i == 1 {
		s := r.String()
		return s[:len(s)-2] + "-" + ds[len(ds)-2:]
	}

	return r.String()
}
