package main

import (
	"sort"
)

func main() {

}

func reorderLogFiles(logs []string) []string {
	let := make([]string, 0)
	dig := make([]string, 0)

	for _, l := range logs {
		c := l[len(l)-1]
		if 'a' <= c && c <= 'z' {
			let = append(let, l)
		} else {
			dig = append(dig, l)
		}
	}

	sort.Slice(let, func(a, b int) bool {
		ai := 0
		for ; ai < len(let[a]); ai++ {
			if let[a][ai] == ' ' {
				break
			}
		}

		bi := 0
		for ; bi < len(let[b]); bi++ {
			if let[b][bi] == ' ' {
				break
			}
		}

		if let[a][ai:] == let[b][bi:] {
			return let[a][:ai] < let[b][:bi]
		}

		return let[a][ai:] < let[b][bi:]
	})

	return append(let[:], dig...)
}
