package main

import "fmt"

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
}

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	i1 := -1
	i2 := -1
	c := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}

		c++
		if c > 2 {
			return false
		}

		if i1 == -1 {
			i1 = i
		} else {
			i2 = i
		}
	}

	if c != 2 {
		return false
	}

	return s1[i1] == s2[i2] && s1[i2] == s2[i1]
}
