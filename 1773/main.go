package main

func main() {
}

var (
	mp = map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
)

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	r := 0
	value := mp[ruleKey]
	for _, item := range items {
		if item[value] == ruleValue {
			r++
		}
	}
	return r
}
