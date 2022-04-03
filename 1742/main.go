package main

import "fmt"

func main() {
	fmt.Println(countBalls(5, 15))
}

func countBalls(lowLimit int, highLimit int) int {
	mp := make([]int, 45)

	fn := lowLimit
	last := 0
	for fn > 0 {
		last += fn % 10
		fn = fn / 10
	}
	mp[last]++
	r := 1

	for i := lowLimit + 1; i <= highLimit; i++ {
		t := 0

		if i%10 != 0 {
			t = last + 1
		} else {
			n := i
			for n > 0 {
				t += n % 10
				n = n / 10
			}
		}

		mp[t]++
		if mp[t] > r {
			r = mp[t]
		}

		last = t
	}

	return r
}
