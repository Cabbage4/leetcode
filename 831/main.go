package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maskPII("LeetCode@LeetCode.com"))
	fmt.Println(maskPII("86-(10)12345678"))
}

func maskPII(s string) string {
	if strings.Contains(s, "@") {
		return email(s)
	}

	return phone(s)
}

func email(s string) string {
	index := strings.Index(s, "@")

	to := func(b byte) string {
		if 'A' <= b && b <= 'Z' {
			return string(b - 'A' + 'a')
		}
		return string(b)
	}

	r := fmt.Sprintf("%s*****%s", to(s[0]), to(s[index-1]))
	for i := index; i < len(s); i++ {
		r += to(s[i])
	}

	return r
}

func phone(s string) string {
	pe := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' || s[i] == '(' || s[i] == ')' || s[i] == ' ' {
			continue
		}
		pe += string(s[i])
	}
	if len(pe) == 10 {
		return "***-***-" + pe[len(pe)-4:]
	} else if len(pe) == 11 {
		return "+*-***-***-" + pe[len(pe)-4:]
	} else if len(pe) == 12 {
		return "+**-***-***-" + pe[len(pe)-4:]
	} else {
		return "+***-***-***-" + pe[len(pe)-4:]
	}
}
