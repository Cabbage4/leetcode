package main

func main() {
}

func halvesAreAlike(s string) bool {
	l := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	mp := make([]bool, 256)
	for _, v := range l {
		mp[v] = true
	}

	c := 0
	for i := 0; i < len(s)/2; i++ {
		if mp[s[i]] {
			c++
		}
	}

	for i := len(s) / 2; i < len(s); i++ {
		if mp[s[i]] {
			c--
		}
	}

	return c == 0
}
