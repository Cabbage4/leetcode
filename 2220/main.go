package main

func main() {
}

func minBitFlips(start int, goal int) int {
	t := start ^ goal
	r := 0
	for t > 0 {
		r += t % 2
		t = t / 2
	}
	return r
}
