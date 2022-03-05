package main

func main() {

}

func minDeletionSize(s []string) int {
	r := 0
	for i := 0; i < len(s[0]); i++ {
		flag := false
		for j := 1; j < len(s); j++ {
			if s[j][i] < s[j-1][i] {
				flag = true
				break
			}
		}

		if flag {
			r++
		}
	}
	return r
}
