package main

import "fmt"

func main() {
	fmt.Println(addStrings("0", "0"))
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("1", "9"))
}

func addStrings(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	result := ""
	offset := 0
	n1I := len(num1) - 1
	n2I := len(num2) - 1
	for n1I >= 0 {
		t := int(18-'9'+int32(num1[n1I]+num2[n2I])-'9') + offset
		result = string('0'+t%10) + result
		offset = t / 10
		n1I--
		n2I--
	}

	if offset > 0 {
		for i := n2I; i >= 0; i-- {
			t := int(9-'9'+int32(num2[i])) + offset
			result = string('0'+t%10) + result
			offset = t / 10
		}

		if offset > 0 {
			result = string('0'+offset) + result
		}
	} else {
		result = num2[0:n2I+1] + result
	}

	return result
}
