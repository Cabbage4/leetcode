package main

import "fmt"

func main() {
	fmt.Println(reformatDate("20th Oct 2052"))
	fmt.Println(reformatDate("6th Jun 1933"))
}

func reformatDate(date string) string {
	m := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	c := 0
	r := ""
	for i := 0; i < len(date); {
		if date[i] == ' ' {
			i++
			continue
		}

		j := i + 1
		for j < len(date) {
			if date[j] != ' ' {
				j++
				continue
			}
			break
		}

		if c == 0 {
			r = date[i : j-2]
			if j-2-i == 1 {
				r = "0" + r
			}
		} else if c == 1 {
			r = m[date[i:j]] + "-" + r
		} else {
			r = date[i:j] + "-" + r
		}

		c++
		i = j
	}

	return r
}
