package main

import "fmt"

func main() {
	//fmt.Println(compareVersion("1.01", "1.0001"))
	fmt.Println(compareVersion("000.00.2", "1.1"))
}

func compareVersion(v1 string, v2 string) int {
	v1l := genList(v1)
	v2l := genList(v2)
	for i := 0; i < len(v1l); i++ {
		var flag int
		if i < len(v2l) {
			flag = compare(v1l[i], v2l[i])
		} else {
			flag = compare(v1l[i], "0")
		}
		if flag != 0 {
			return flag
		}
	}

	for i := len(v1l); i < len(v2l); i++ {
		var flag int
		flag = compare("0", v2l[i])
		if flag != 0 {
			return flag
		}
	}

	return 0
}

func genList(v1 string) []string {
	r := make([]string, 0)
	for i := 0; i < len(v1); {
		if v1[i] == '.' {
			i++
			continue
		}

		for i < len(v1) && v1[i] == '0' {
			i++
		}

		if i == len(v1) {
			r = append(r, "0")
			break
		}

		if v1[i] == '.' {
			r = append(r, "0")
			continue
		}

		j := i + 1
		for j < len(v1) {
			if v1[j] == '.' {
				break
			}
			j++
		}

		r = append(r, v1[i:j])
		i = j
	}

	return r
}

func compare(a, b string) int {
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	if a < b {
		return -1
	} else if a > b {
		return 1
	}

	return 0
}
