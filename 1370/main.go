package main

import "fmt"

func main() {
	fmt.Println(sortString("aaaabbbbcccc") == "abccbaabccba")
	fmt.Println(sortString("rat") == "art")
}

func sortString(s string) string {
	r := make([]byte, 0)

	mp := make([]int, 26)
	for _, v := range s {
		t := int(v - 'a')
		mp[t]++
	}

	for len(s) > len(r) {
		for i := 0; i < 26; i++ {
			if mp[i] != 0 {
				mp[i]--
				r = append(r, byte(i+'a'))
			}
		}

		for i := 25; i >= 0; i-- {
			if mp[i] != 0 {
				mp[i]--
				r = append(r, byte(i+'a'))
			}
		}
	}

	return string(r)
}
