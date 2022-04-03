package main

func main() {
}

func largestAltitude(gain []int) int {
	r := 0
	c := 0
	for _, v := range gain {
		c += v
		if c > r {
			r = c
		}
	}
	return r
}
