package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

func subdomainVisits(cpdomains []string) []string {
	rmp := make(map[string]int)

	parse := func(str string) {
		infos := strings.Split(str, " ")
		count, _ := strconv.Atoi(infos[0])
		name := infos[1]

		for i := len(name) - 1; i >= 0; {
			j := i - 1
			for j >= 0 && name[j] != '.' {
				j--
			}

			if j == -1 {
				rmp[name] += count
			} else {
				rmp[name[j+1:]] += count
			}

			i = j
		}
	}

	for _, domain := range cpdomains {
		parse(domain)
	}

	var r []string
	for k, v := range rmp {
		r = append(r, fmt.Sprintf("%d %s", v, k))
	}
	return r
}
