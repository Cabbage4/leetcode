package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-wg", 4))
}

func licenseKeyFormatting(s string, k int) string {
	ns := strings.ToUpper(strings.ReplaceAll(s, "-", ""))
	if len(ns) == 0 {
		return ""
	}

	r := make([]byte, 0)
	i := len(ns)
	for i >= k {
		r = append([]byte("-"+ns[i-k:i]), r...)
		i = i - k
	}
	if i > 0 {
		return ns[:i] + string(r)
	}
	return string(r[1:])
}
