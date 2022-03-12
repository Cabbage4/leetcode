package main

import "fmt"

func main() {
	fmt.Println(modifyString("z?z"))
	fmt.Println(modifyString("?zs"))
	fmt.Println(modifyString("ubv?w"))
}

func modifyString(s string) string {
	if len(s) == 1 {
		if s[0] == '?' {
			return "a"
		}
		return s
	}

	r := []byte(s)

	for i := 0; i < len(r); i++ {
		if r[i] != '?' {
			continue
		}

		if i == 0 {
			if r[i+1] == 'a' {
				r[i] = 'z'
			} else {
				r[i] = 'a'
			}
			continue
		}

		if i+1 == len(r) || r[i+1] == '?' {
			if r[i-1] == 'a' {
				r[i] = 'z'
			} else {
				r[i] = 'a'
			}
			continue
		}

		r[i] = pick(r[i+1], r[i-1])
	}

	return string(r)
}

func pick(a, b byte) byte {
	if a > b {
		a, b = b, a
	}

	if a == b {
		if a == 'z' {
			return 'a'
		}
		return 'z'
	}

	if b-a > 1 {
		return a + 1
	}

	if a == 'a' {
		return 'z'
	}

	return 'a'
}
