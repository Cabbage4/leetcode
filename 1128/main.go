package main

func main() {

}

func numEquivDominoPairs(dominoes [][]int) int {
	r := 0
	mp := make([]int, 100)

	for i := 0; i < len(dominoes); i++ {
		a, b := ad(dominoes[i][0], dominoes[i][1])
		mp[a*10+b]++
	}

	for _, v := range mp {
		if v <= 1 {
			continue
		}
		r += v * (v - 1) / 2
	}

	return r
}

func ad(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
