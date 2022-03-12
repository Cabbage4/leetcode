package main

func main() {
}

func maxPower(s string) int {
	r := 0
	for i := 0; i < len(s); {
		j := i + 1
		for j < len(s) && s[i] == s[j] {
			j++
		}

		if j-i > r {
			r = j - i
		}

		i = j
	}
	return r
}
