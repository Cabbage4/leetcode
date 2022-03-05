package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}

func numUniqueEmails(emails []string) int {
	r := 0
	mp := make(map[string]bool)

	for _, e := range emails {
		cur := bytes.NewBuffer(make([]byte, 0))
		withPlus := false
		for i := 0; i < len(e); i++ {
			if e[i] == '@' {
				cur.WriteString(e[i:])
				break
			}

			if withPlus {
				continue
			}

			if e[i] == '+' {
				withPlus = true
				continue
			}

			if e[i] != '.' {
				cur.WriteByte(e[i])
			}
		}

		s := cur.String()
		if !mp[s] {
			mp[s] = true
			r++
		}
	}

	return r
}
