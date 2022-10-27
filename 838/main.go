package main

import "fmt"

func main() {
	fmt.Println(pushDominoes(".L.R...LR..L.."))
	fmt.Println(pushDominoes("RR.L"))
}

func pushDominoes(dominoes string) string {
	r := []byte(dominoes)
	var left byte = 'L'
	for i := 0; i < len(dominoes); {
		j := i
		for j < len(dominoes) && dominoes[j] == '.' {
			j++
		}

		var right byte = 'R'
		if j < len(dominoes) {
			right = dominoes[j]
		}

		if left == right {
			ii := i
			for ii < j {
				r[ii] = left
				ii++
			}
		} else if left == 'R' && right == 'L' {
			ii := i
			jj := j - 1
			for ii < jj {
				r[ii] = 'R'
				r[jj] = 'L'
				ii++
				jj--
			}
		}

		left = right
		i = j + 1
	}
	return string(r)
}
