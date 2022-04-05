package main

func main() {
}

func maximumPopulation(logs [][]int) int {
	list := make([]int, 101)
	for _, v := range logs {
		list[v[0]-1950]++
		list[v[1]-1950]--
	}

	r := 0
	c := -1
	for i := 1; i <= 100; i++ {
		list[i] += list[i-1]
		if list[i] > c {
			c = list[i]
			r = i + 1950
		}
	}

	return r
}
