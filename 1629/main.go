package main

import "fmt"

/**
[9,29,49,50]
"cbcd"
*/
func main() {
	fmt.Println(slowestKey([]int{9, 29, 49, 50}, "cbcd"))
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	r := keysPressed[0]
	t := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		c := releaseTimes[i] - releaseTimes[i-1]
		if c > t || (c == t && keysPressed[i] > r) {
			t = c
			r = keysPressed[i]
		}
	}

	return r
}
