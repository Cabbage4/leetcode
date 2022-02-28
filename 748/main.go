package main

import "fmt"

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
}

var prime = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

func shortestCompletingWord(licensePlate string, words []string) string {
	license := product(licensePlate)

	result := ""
	for _, w := range words {
		if len(w) < len(result) || result == "" {
			if product(w)%license == 0 {
				result = w
			}
		}
	}

	return result
}

func product(licensePlate string) int64 {
	var result int64 = 1
	for _, v := range licensePlate {
		if 'a' <= v && v <= 'z' {
			result *= int64(prime[v-'a'])
		}
		if 'A' <= v && v <= 'Z' {
			result *= int64(prime[v-'A'])
		}
	}
	return result
}
