package main

func main() {
}

func countGoodRectangles(rectangles [][]int) int {
	mp := make(map[int]int)
	i := 0
	for _, v := range rectangles {
		t := v[0]
		if v[0] > v[1] {
			t = v[1]
		}
		mp[i]++
		if t > i {
			i = t
		}
	}

	return mp[i]
}
