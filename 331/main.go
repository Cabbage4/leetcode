package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidSerialization("1,#") == false)
	fmt.Println(isValidSerialization("9,#,1,#,#") == true)
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#") == true)
}

func isValidSerialization(preorder string) bool {
	slot := 1
	for i := 0; i < len(preorder); {
		if slot == 0 {
			return false
		}

		if preorder[i] == ',' {
			i++
			continue
		}

		j := i + 1
		for j < len(preorder) && preorder[j] != ',' {
			j++
		}

		s := preorder[i:j]
		if s == "#" {
			slot--
		} else {
			slot++
		}

		i = j
	}

	return slot == 0
}
