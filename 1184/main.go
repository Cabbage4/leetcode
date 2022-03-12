package main

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	cur := 0
	for i := start; i != destination; {
		cur += distance[i]
		i = (i + 1) % len(distance)
	}

	r := cur
	cur = 0

	for i := (start - 1 + len(distance)) % len(distance); ; {
		cur += distance[i]
		if i == destination {
			break
		}
		i = (i - 1 + len(distance)) % len(distance)
	}
	if cur < r {
		r = cur
	}

	return r
}
