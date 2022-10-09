package main

import "fmt"

func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
}

func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}

	si := 0
	ei := 0

	for si < len(start) && ei < len(end) {
		for si < len(start) && start[si] == 'X' {
			si++
		}
		for ei < len(end) && end[ei] == 'X' {
			ei++
		}

		if ei < len(end) && si < len(start) {
			if end[ei] != start[si] {
				return false
			}

			c := end[ei]
			if c == 'L' && si < ei || c == 'R' && ei < si {
				return false
			}

			si++
			ei++
		}
	}

	for si < len(start) {
		if start[si] != 'X' {
			return false
		}
		si++
	}

	for ei < len(end) {
		if end[ei] != 'X' {
			return false
		}
		ei++
	}

	return true
}
