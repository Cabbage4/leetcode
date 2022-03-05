package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}))
}

func lemonadeChange(bills []int) bool {
	a := 0
	b := 0
	c := 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			a++
			continue
		}

		if bills[i] == 10 {
			if a <= 0 {
				return false
			}

			a--
			b++
			continue
		}

		if a < 1 || (b < 1 && a < 3) {
			return false
		}

		if b < 1 {
			a = a - 3
		} else {
			a--
			b--
		}

		c++
	}

	return true
}
