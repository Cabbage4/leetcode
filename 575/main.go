package main

import "math"

func main() {

}

func distributeCandies(candyType []int) int {
	ln := int(math.Pow10(5))
	mp := make([]bool, ln*2+1)
	count := 0
	for _, v := range candyType {
		if !mp[v+ln] {
			mp[v+ln] = true
			count++
		}
	}
	if len(candyType)/2 <= count {
		return len(candyType) / 2
	} else {
		return count
	}
}
