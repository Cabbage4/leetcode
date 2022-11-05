package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(makeIntegerBeautiful(204932336, 16))
	fmt.Println(makeIntegerBeautiful(30978, 6))
	fmt.Println(makeIntegerBeautiful(839, 11))
	fmt.Println(makeIntegerBeautiful(409, 9))
	fmt.Println(makeIntegerBeautiful(8, 2))
	fmt.Println(makeIntegerBeautiful(16, 6))
	fmt.Println(makeIntegerBeautiful(1, 1))
	fmt.Println(makeIntegerBeautiful(467, 6))
}

func makeIntegerBeautiful(n int64, target int) int64 {
	ns := []byte(strconv.FormatInt(n, 10))
	if count(n) <= int64(target) {
		return 0
	}

	last := byte('0' + n%10)
	nln := ln(n)
	for i := nln - 2; i >= 0; i-- {
		origin := ns[i]
		if ns[i] == '9' {
			ns[i+1] = '0'
			last = origin
			continue
		}

		ns[i+1] = '0'
		ns[i] = ns[i] + 1
		vh, _ := strconv.ParseInt(string(ns), 10, 64)
		if count(vh) > int64(target) {
			last = origin
			continue
		}

		if last == '9' {
			ns[i+1] = last
		} else {
			ns[i+1] = last + 1
		}
		ns[i] = origin
		vl, _ := strconv.ParseInt(string(ns), 10, 64)
		for j := vl; j <= vh; j++ {
			if count(j) <= int64(target) {
				return j - n
			}
		}
	}

	return int64(math.Pow10(nln)) - n
}

func count(n int64) int64 {
	var r int64
	for n != 0 {
		r += n % 10
		n = n / 10
	}
	return r
}

func ln(n int64) int {
	var r int
	for n != 0 {
		r++
		n = n / 10
	}
	return r
}
