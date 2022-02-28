package main

import "strings"

func main() {

}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	return strings.Contains(goal+goal, s)
}
