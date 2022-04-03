package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(interpret("G()(al)"))
	fmt.Println(interpret("G()()()()(al)"))
}

func interpret(command string) string {
	r := bytes.NewBuffer(make([]byte, 0))

	flag := false
	fi := -1
	for i := 0; i < len(command); i++ {
		if command[i] == '(' {
			flag = true
			fi = i
			continue
		}

		if command[i] == ')' {
			flag = false
			if i-fi == 1 {
				r.WriteByte('o')
			} else {
				r.WriteString(command[fi+1 : i])
			}

			continue
		}

		if !flag {
			r.WriteByte(command[i])
			continue
		}
	}

	return r.String()
}
