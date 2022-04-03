package main

func main() {
}

func minOperations(logs []string) int {
	r := 1
	for i := 0; i < len(logs); i++ {
		if logs[i] == "./" {
			continue
		}
		if logs[i] == "../" {
			if r == 1 {
				continue
			}

			r--
			continue
		}

		r++
	}

	return r - 1
}
