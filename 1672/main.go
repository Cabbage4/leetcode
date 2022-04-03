package main

func main() {
}

func maximumWealth(accounts [][]int) int {
	r := 0
	for i := range accounts {
		tr := 0
		for j := range accounts[i] {
			tr += accounts[i][j]
		}
		if tr > r {
			r = tr
		}
	}
	return r
}
