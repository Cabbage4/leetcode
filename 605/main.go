package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	for i := range flowerbed {
		cl := i == 0 || flowerbed[i-1] == 0
		cr := i == len(flowerbed)-1 || flowerbed[i+1] == 0
		if cl && cr && flowerbed[i] == 0 {
			count++
			flowerbed[i] = 1
			if count >= n {
				return true
			}
		}
	}

	return count >= n
}
