package main

func main() {
}

func numberOfMatches(n int) int {
	r := 0
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
			r += n
		} else {
			n = n/2 + 1
			r += n - 1
		}
	}
	return r
}
