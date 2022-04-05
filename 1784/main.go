package main

import "strings"

func main() {
}

func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
