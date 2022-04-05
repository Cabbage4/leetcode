package main

import "fmt"

func main() {
	fmt.Println(replaceDigits("a1c1e1"))
}

func replaceDigits(s string) string {
	r := []byte(s)
	for i := 1; i < len(s); i = i + 2 {
		r[i] = r[i-1] + r[i] - '0'
	}

	return string(r)
}
