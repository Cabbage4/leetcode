package main

func main() {
}

func makeFancyString(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) < 2 {
			stack = append(stack, s[i])
			continue
		}

		if stack[len(stack)-1] == s[i] &&
			stack[len(stack)-2] == s[i] {
			continue
		}

		stack = append(stack, s[i])
	}

	return string(stack)
}
