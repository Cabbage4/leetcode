package main

import "fmt"

func main() {
	m := Constructor()
	m.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(m.Search("hlhlo"))
}

type MagicDictionary struct {
	mp map[string][26]int
}

func Constructor() MagicDictionary {
	return MagicDictionary{mp: make(map[string][26]int)}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	for _, d := range dictionary {
		var c [26]int
		for i := 0; i < len(d); i++ {
			c[d[i]-'a']++
		}
		m.mp[d] = c
	}
}

func (m *MagicDictionary) Search(sw string) bool {
	var c [26]int
	for i := 0; i < len(sw); i++ {
		c[sw[i]-'a']++
	}

	for k, cv := range m.mp {
		if len(sw) != len(k) {
			continue
		}

		count := 0
		flag := false
		for i := 0; i < 26; i++ {
			count += abs(cv[i], c[i])

			if count > 2 {
				flag = true
				break
			}
		}

		if flag || count == 0 {
			continue
		}

		count = 0
		for i := 0; i < len(k); i++ {
			if k[i] == sw[i] {
				continue
			}
			count++
			if count > 1 {
				flag = true
				break
			}
		}

		if flag || count == 0 {
			continue
		}

		return true
	}

	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
