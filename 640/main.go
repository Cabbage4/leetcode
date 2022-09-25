package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(solveEquation("0x=0"))
	fmt.Println(solveEquation("3x=33+22+11"))
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
	fmt.Println(solveEquation("x=x"))
	fmt.Println(solveEquation("2x=x"))
}

func solveEquation(e string) string {
	sign := 1
	value := 0
	factor := 0
	for i := 0; i < len(e); {
		if e[i] == '=' {
			sign = -1
			i++
			continue
		}

		curFactor := 1
		if e[i] == '+' {
			i++
		} else if e[i] == '-' {
			i++
			curFactor = -1
		}

		v := 0
		hasDigit := false
		for i < len(e) && unicode.IsDigit(rune(e[i])) {
			v = v*10 + int(e[i]-'0')
			hasDigit = true
			i++
		}

		if i < len(e) && e[i] == 'x' {
			if hasDigit {
				factor += curFactor * v * sign
			} else {
				factor += 1 * sign * curFactor
			}

			i++
			continue
		}

		value += v * sign * curFactor
	}

	if factor == 0 {
		if value == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}

	return fmt.Sprintf("x=%d", -value/factor)
}
