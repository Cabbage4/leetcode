package main

func main() {
}

func largestOddNumber(num string) string {
	for num != "" {
		c := num[len(num)-1] - '0'
		if c%2 != 0 {
			return num
		}
		num = num[:len(num)-1]
	}

	return ""
}
