package main

import "fmt"

func main() {
	fmt.Println(customSortString("exv", "xwvee"))
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("kqep", "pekeq"))
}

func customSortString(order string, s string) string {
	var mp [26]int
	for i := 0; i < len(order); i++ {
		mp[order[i]-'a'] = i + 1
	}

	sb := []byte(s)
	list := make([]int, 0)
	for i := 0; i < len(s); i++ {
		o := mp[s[i]-'a']
		if o == 0 {
			continue
		}

		list = append(list, i)

		if len(list) == 0 {
			continue
		}

		tmpIndex := i
		listIndex := len(list) - 2
		for 0 <= listIndex && o < mp[sb[list[listIndex]]-'a'] {
			sb[tmpIndex], sb[list[listIndex]] = sb[list[listIndex]], sb[tmpIndex]
			tmpIndex = list[listIndex]
			listIndex--
		}
	}
	return string(sb)
}
