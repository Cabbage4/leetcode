package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := []string{"01:15", "02:00"}
	b := []string{"02:01", "03:00"}
	fmt.Println(haveConflict(a, b))
}

func haveConflict(event1 []string, event2 []string) bool {
	parse := func(event string) int {
		infos := strings.Split(event, ":")
		v1, _ := strconv.Atoi(infos[0])
		v2, _ := strconv.Atoi(infos[1])
		return v1*60 + v2
	}

	a1, a2 := parse(event1[0]), parse(event1[1])
	b1, b2 := parse(event2[0]), parse(event2[1])
	if !(a1 > b2 || a2 < b1) {
		return true
	}
	return false
}
