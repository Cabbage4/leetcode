package main

import "fmt"

func main() {
	fmt.Println(expressiveWords("dddiiiinnssssssoooo", []string{"dinnssoo", "ddinso", "ddiinnso", "ddiinnssoo", "ddiinso", "dinsoo", "ddiinsso", "dinssoo", "dinso"}))
	fmt.Println(expressiveWords("zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}))
	fmt.Println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
}

func expressiveWords(s string, words []string) int {
	genInfo := func(str string) [][]int {
		r := make([][]int, 0)
		for i := 0; i < len(str); {
			j := i + 1
			for j < len(str) && str[i] == str[j] {
				j++
			}

			r = append(r, []int{int(str[i]), j - i})

			i = j
		}
		return r
	}

	var r int
	sInfo := genInfo(s)
	for _, w := range words {
		wInfo := genInfo(w)
		if len(wInfo) != len(sInfo) {
			continue
		}

		flag := true
		for i := 0; i < len(wInfo); i++ {
			if wInfo[i][0] != sInfo[i][0] {
				flag = false
				break
			}

			if wInfo[i][1] == sInfo[i][1] {
				continue
			}

			if sInfo[i][1] < wInfo[i][1] {
				flag = false
				break
			}

			if sInfo[i][1] <= 2 {
				flag = false
				break
			}
		}

		if flag {
			r++
		}
	}

	return r
}
