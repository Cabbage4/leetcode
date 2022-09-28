package main

func main() {
}

func findRedundantConnection(edges [][]int) []int {
	mp := make([]int, len(edges)+1)
	for i := range mp {
		mp[i] = i
	}

	var find func(int) int
	find = func(i int) int {
		if mp[i] == i {
			return i
		}
		mp[i] = find(mp[i])
		return mp[i]
	}

	union := func(i, j int) bool {
		ii, jj := find(i), find(j)
		if ii == jj {
			return false
		}
		mp[ii] = jj
		return true
	}

	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}

	return nil
}
