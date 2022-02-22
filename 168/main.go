package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(28) == "AB")
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(52) == "AZ")
}

func convertToTitle(columnNumber int) string {
	mp := make(map[int]string)
	for i := 0; i <= 26; i++ {
		mp[i+1] = string(rune('A' + i))
	}

	if columnNumber <= 26 {
		return mp[columnNumber]
	}

	res := ""
	nv := columnNumber
	for nv >= 26 {
		if nv%26 == 0 {
			res = "Z" + res
			nv = nv/26 - 1
			continue
		}

		res = mp[nv%26] + res
		nv = nv / 26
	}

	if nv != 0 {
		res = mp[nv%26] + res
	}

	return res
}
