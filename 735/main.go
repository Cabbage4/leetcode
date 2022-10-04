package main

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}))
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}

func asteroidCollision(asteroids []int) []int {
	var r []int

	for i := 0; i < len(asteroids); i++ {
		if len(r) > 0 && r[len(r)-1] > 0 && asteroids[i] < 0 {
			for len(r) > 0 && r[len(r)-1] < -asteroids[i] && r[len(r)-1] > 0 {
				r = r[:len(r)-1]
			}

			if len(r) > 0 && r[len(r)-1]+asteroids[i] == 0 {
				r = r[:len(r)-1]
				continue
			}

			if len(r) > 0 && r[len(r)-1] > -asteroids[i] {
				continue
			}

			r = append(r, asteroids[i])
			continue
		}

		r = append(r, asteroids[i])
	}

	return r
}
