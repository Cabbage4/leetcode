package main

import "fmt"

func main() {
	r := Constructor()
	fmt.Println(r.Ping(1))
	fmt.Println(r.Ping(100))
	fmt.Println(r.Ping(3001))
	fmt.Println(r.Ping(3002))
}

type RecentCounter struct {
	list []int
}

func Constructor() RecentCounter {
	return RecentCounter{list: make([]int, 0)}
}

func (r *RecentCounter) Ping(t int) int {
	r.list = append(r.list, t)
	if t-3000 <= 0 {
		return len(r.list)
	}

	index := binarySearch(r.list, t-3000)
	result := len(r.list) - index
	r.list = r.list[index:]
	return result
}

func binarySearch(l []int, target int) int {
	left := 0
	right := len(l)
	for left < right {
		mid := (left + right) / 2
		if l[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
