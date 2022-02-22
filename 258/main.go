package main

func main() {

}

func addDigits(num int) int {
	if num >= 0 && num <= 9 {
		return num
	}

	next := 0
	for num != 0 {
		next += num % 10
		num = num / 10

		if num == 0 && next > 9 {
			num = next
			next = 0
		}
	}

	return next
}
