package main

func main() {

}

func numberOfLines(widths []int, s string) []int {
	total := 0
	cur := 0
	for _, v := range s {
		if cur+widths[v-'a'] > 100 {
			total++
			cur = widths[v-'a']
			continue
		}
		cur += widths[v-'a']
	}

	if cur != 0 {
		total++
	}
	return []int{total, cur}
}
