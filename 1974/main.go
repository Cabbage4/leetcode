package main

import "fmt"

func main() {
	fmt.Println(minTimeToType("abc"))
	fmt.Println(minTimeToType("bza"))
}

func minTimeToType(word string) int {
	r := h('a', word[0]) + 1
	for i := 1; i < len(word); i++ {
		r += 1 + h(word[i-1], word[i])
	}
	return r
}

func h(a, b byte) int {
	t1 := int((a - b + 26) % 26)
	t2 := int((b - a + 26) % 26)

	if t1 > t2 {
		return t2
	}
	return t1
}
