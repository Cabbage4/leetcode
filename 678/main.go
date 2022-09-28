package main

import "fmt"

func main() {
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*)))"))
}

func checkValidString(s string) bool {
	list := make([]int, 0)
	cache := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			list = append(list, i)
		} else if s[i] == '*' {
			cache = append(cache, i)
		} else {
			if len(list) == 0 {
				if len(cache) == 0 {
					return false
				}
				cache = cache[:len(cache)-1]
			} else {
				list = list[:len(list)-1]
			}
		}
	}

	if len(list) == 0 {
		return true
	}

	if len(list) > len(cache) {
		return false
	}

	for i := 0; i < len(list); i++ {
		for len(cache) > 0 && cache[0] < list[i] {
			cache = cache[1:]
		}

		if len(cache) == 0 {
			return false
		}
		cache = cache[1:]
	}
	return true
}
