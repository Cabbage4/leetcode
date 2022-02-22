package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	res := ""

	var with uint8 = 0
	ai := len(a) - 1
	bi := len(b) - 1
	for ai >= 0 && bi >= 0 {
		n := (a[ai] - '0') + (b[bi] - '0') + with
		if n > 1 {
			with = 1
		} else {
			with = 0
		}

		res = string(n%2+'0') + res
		ai--
		bi--
	}

	for ai >= 0 {
		n := a[ai] - '0' + with
		if n > 1 {
			with = 1
		} else {
			with = 0
		}
		res = string(n%2+'0') + res
		ai--
	}

	for bi >= 0 {
		n := b[bi] - '0' + with
		if n > 1 {
			with = 1
		} else {
			with = 0
		}
		res = string(n%2+'0') + res
		bi--
	}

	if with == 1 {
		res = "1" + res
	}

	return res
}
