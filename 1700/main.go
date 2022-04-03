package main

func main() {
}

func countStudents(students []int, sandwiches []int) int {
	o := 0
	e := 0
	for _, v := range students {
		if v == 1 {
			e++
		} else {
			o++
		}
	}

	for i, s := range sandwiches {
		if (s == 0 && o == 0) || (s == 1 && e == 0) {
			return len(sandwiches) - i
		}

		if s == 1 {
			e--
		} else {
			o--
		}
	}

	return 0
}
