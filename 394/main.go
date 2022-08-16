package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("a2[c]"))
}

func decodeString(s string) string {
	r := bytes.NewBufferString("")
	for i := 0; i < len(s); {
		if 'a' <= s[i] && s[i] <= 'z' {
			r.WriteByte(s[i])
			i++
			continue
		}

		j := i
		for j < len(s) && '0' <= s[j] && s[j] <= '9' {
			j++
		}

		count, _ := strconv.Atoi(s[i:j])
		left := j + 1
		right := left
		for cc := 1; cc > 0; {
			if s[right] == ']' {
				cc--
			} else if s[right] == '[' {
				cc++
			}

			right++
		}

		r.WriteString(strings.Repeat(decodeString(s[left:right-1]), count))
		i = right
	}

	return r.String()
}
