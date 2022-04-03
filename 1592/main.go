package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
}

func reorderSpaces(text string) string {
	sc := 0
	wl := make([]string, 0)
	for i := 0; i < len(text); {
		if text[i] == ' ' {
			sc++
			i++
			continue
		}

		j := i + 1
		for ; j < len(text) && text[j] != ' '; j++ {
		}
		wl = append(wl, text[i:j])

		i = j
	}

	ll := len(wl) - 1
	if ll == 0 {
		return wl[0] + strings.Repeat(" ", sc)
	}

	return strings.Join(wl, strings.Repeat(" ", sc/ll)) + strings.Repeat(" ", sc%ll)
}
