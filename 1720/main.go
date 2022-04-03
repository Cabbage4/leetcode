package main

func main() {
}

func decode(encoded []int, first int) []int {
	r := make([]int, len(encoded)+1)
	r[0] = first
	for i := 1; i < len(r); i++ {
		r[i] = encoded[i-1] ^ r[i-1]
	}
	return r
}
