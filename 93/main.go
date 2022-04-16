package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses("101023"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("25525511135"))
}

func restoreIpAddresses(s string) []string {
	r := make([]string, 0)
	dfs(s, 0, &r, make([]string, 0))
	return r
}

func dfs(s string, index int, r *[]string, cr []string) {
	if len(cr) == 4 && index == len(s) {
		*r = append(*r, strings.Join(cr, "."))
		return
	}

	if index < len(s) && s[index] == '0' {
		dfs(s, index+1, r, append(cr, "0"))
		return
	}

	for i := index; i < len(s) && i < index+3; i++ {
		t, _ := strconv.ParseInt(s[index:i+1], 10, 64)
		if t > 255 {
			break
		}

		dfs(s, i+1, r, append(cr, s[index:i+1]))
	}
}
