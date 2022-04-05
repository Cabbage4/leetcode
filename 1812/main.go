package main

import "fmt"

func main() {
	fmt.Println(squareIsWhite("a1"))
}

func squareIsWhite(cs string) bool {
	w1 := cs[0] - 'a'
	w2 := cs[1] - '1'
	if w1%2 == 0 {
		return w2%2 != 0
	} else {
		return w2%2 == 0
	}
}
