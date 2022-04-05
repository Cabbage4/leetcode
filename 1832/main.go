package main

func main() {
}

func checkIfPangram(sentence string) bool {
	bit := 0
	for i := range sentence {
		bit = bit | (1 << (sentence[i] - 'a'))
	}
	return bit == ((1 << 26) - 1)
}
