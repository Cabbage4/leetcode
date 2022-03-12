package main

import "fmt"

func main() {
	fmt.Println(isPathCrossing("NESWW"))
}

func isPathCrossing(path string) bool {
	mp := make(map[string]bool)
	mp["0:0"] = true

	x := 0
	y := 0
	for _, p := range path {
		switch p {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}

		key := fmt.Sprintf("%d:%d", x, y)
		if mp[key] {
			return true
		}
		mp[key] = true
	}

	return false
}
