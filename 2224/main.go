package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertTime("09:41", "10:34"))
	fmt.Println(convertTime("11:00", "11:01"))
	fmt.Println(convertTime("02:30", "04:35"))
}

func convertTime(current string, correct string) int {
	oneH, _ := strconv.Atoi(current[0:2])
	twoH, _ := strconv.Atoi(correct[0:2])

	r := twoH - oneH

	oneM, _ := strconv.Atoi(current[3:])
	twoM, _ := strconv.Atoi(correct[3:])
	d := (twoM - oneM + 60) % 60

	if oneM > twoM {
		r--
	}

	r += d/15 + (d%15)/5 + (d%15)%5

	return r
}
