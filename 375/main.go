package main

import "fmt"

func main() {
	fmt.Println(getMoneyAmount(10))
	//fmt.Println(getMoneyAmount(2))
}

func getMoneyAmount(n int) int {
	mp := make([][]int, n+2)
	for i := range mp {
		mp[i] = make([]int, n+2)
		for j := range mp[i] {
			mp[i][j] = -1
		}
	}

	return h(1, n, mp)
}

func h(left, right int, mp [][]int) int {
	r := mp[left][right]
	if r != -1 {
		return r
	}

	ln := right - left + 1
	if ln == 1 || ln == 0 {
		return 0
	}

	if ln == 2 {
		return left
	}

	min := left + h(left+1, right, mp)
	for value := left + 1; value <= right; value++ {
		l := value + h(left, value-1, mp)
		r := value + h(value+1, right, mp)

		tmp := r
		if tmp < l {
			tmp = l
		}

		if tmp < min {
			min = tmp
		}
	}

	mp[left][right] = min
	return min
}
