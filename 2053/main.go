package main

func main() {
}

func kthDistinct(arr []string, k int) string {
	mp := make(map[string]int)
	for _, s := range arr {
		mp[s]++
	}

	for _, s := range arr {
		if mp[s] == 1 {
			k--
		}
		if k == 0 {
			return s
		}
	}

	return ""
}
