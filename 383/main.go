package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	mp := make([]int, 256)
	for _, v := range magazine {
		mp[v]++
	}

	for _, v := range ransomNote {
		mp[v]--
		if mp[v] < 0 {
			return false
		}
	}

	return true
}
