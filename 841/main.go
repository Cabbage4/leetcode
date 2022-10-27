package main

import "fmt"

func main() {
	data := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(data))
}

func canVisitAllRooms(rooms [][]int) bool {
	targetN := len(rooms)
	book := make(map[int]bool)
	book[0] = true

	q := []int{0}
	for len(q) != 0 {
		t := q[0]
		q = q[1:]

		for _, r := range rooms[t] {
			if book[r] {
				continue
			}

			q = append(q, r)
			book[r] = true
		}
	}

	return len(book) == targetN
}
