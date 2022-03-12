package main

import "fmt"

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}

func restoreString(s string, indices []int) string {
	r := []byte(s)

	for i, v := range indices {
		r[v] = s[i]
	}

	return string(r)
}
