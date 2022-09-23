package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("abc", "bbbca"))
	fmt.Println(checkInclusion("adc", "dcda"))
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var mp [26]int
	for i := range s1 {
		mp[s1[i]-'a']++
		mp[s2[i]-'a']--
	}

	var diff int
	for _, v := range mp {
		if v != 0 {
			diff++
		}
	}

	if diff == 0 {
		return true
	}

	for i := 1; i < len(s2)-len(s1)+1; i++ {
		x := s2[i-1] - 'a'
		y := s2[i+len(s1)-1] - 'a'
		if x == y {
			continue
		}

		if mp[x] == 0 {
			diff++
		}
		mp[x]++
		if mp[x] == 0 {
			diff--
		}

		if mp[y] == 0 {
			diff++
		}
		mp[y]--
		if mp[y] == 0 {
			diff--
		}

		if diff == 0 {
			return true
		}
	}

	return false
}
