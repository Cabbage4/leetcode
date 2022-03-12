package main

func main() {
}

func subtractProductAndSum(n int) int {
	p := 1
	s := 0
	for n > 0 {
		t := n % 10
		p *= t
		s += t
		n = n / 10
	}

	return p - s
}
