package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(fractionAddition("-1/2+1/3"))
	fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	fmt.Println(fractionAddition("1/3-1/2"))
	fmt.Println(fractionAddition("-1/2+1/2"))
}

func fractionAddition(e string) string {
	list := parse(e)
	up := list[0][0]
	down := list[0][1]
	for _, p := range list[1:] {
		newDown := lcm(down, p[1])
		newUp := newDown/down*up + newDown/p[1]*p[0]
		if newUp == 0 {
			up = 0
			down = 1
			continue
		}

		c := gcd(newDown, newUp)
		up = newUp / c
		down = newDown / c
	}

	if up > 0 && down < 0 {
		return fmt.Sprintf("-%d/%d", up, -down)
	}

	return fmt.Sprintf("%d/%d", up, down)
}

func parse(e string) [][]int {
	var r [][]int

	for i := 0; i < len(e); {
		j := i + 1
		for j < len(e) && e[j] != '-' && e[j] != '+' {
			j++
		}

		upDown := strings.Split(e[i:j], "/")
		a, _ := strconv.Atoi(upDown[0])
		b, _ := strconv.Atoi(upDown[1])
		r = append(r, []int{a, b})

		i = j
	}

	return r
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
