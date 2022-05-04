package main

import "fmt"

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
}

func getHint(secret string, guess string) string {
	mp := make([]int, 256)
	t1, t2 := 0, 0
	for i := 0; i < len(secret); i++ {
		if guess[i] == secret[i] {
			t1++
		} else {
			mp[secret[i]]++
		}
	}

	for i := 0; i < len(guess); i++ {
		if guess[i] != secret[i] {
			if mp[guess[i]] != 0 {
				mp[guess[i]]--
				t2++
			}
		}
	}

	return fmt.Sprintf("%dA%dB", t1, t2)
}
