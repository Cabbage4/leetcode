package main

import "fmt"

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
}

func findOcurrences(text string, first string, second string) []string {
	list := make([]string, 0)
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			continue
		}

		j := i
		for j < len(text) {
			if text[j] != ' ' {
				j++
				continue
			}
			break
		}

		list = append(list, text[i:j])

		i = j
	}

	r := make([]string, 0)
	if len(list) < 3 {
		return r
	}

	for i := 2; i < len(list); i++ {
		if list[i-2] == first && list[i-1] == second {
			r = append(r, list[i])
		}
	}

	return r
}
