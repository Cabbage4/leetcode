package main

func main() {

}

func checkRecord(s string) bool {
	ac := 0
	count := 0
	for _, v := range s {
		if v == 'A' {
			ac++
			if ac >= 2 {
				return false
			}
		}

		if v == 'L' {
			count++
			if count >= 3 {
				return false
			}
			continue
		}

		count = 0
	}
	return true
}
