package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestNiceSubarray([]int{135745088, 609245787, 16, 2048, 2097152}))
	fmt.Println(longestNiceSubarray([]int{84139415, 693324769, 614626365, 497710833, 615598711, 264, 65552, 50331652, 1, 1048576, 16384, 544, 270532608, 151813349, 221976871, 678178917, 845710321, 751376227, 331656525, 739558112, 267703680}))
	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13}))
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10}))
	fmt.Println(longestNiceSubarray([]int{904163577, 321202512, 470948612, 490925389, 550193477, 87742556, 151890632, 655280661, 4, 263168, 32, 573703555, 886743681, 937599702, 120293650, 725712231, 257119393}))
}

func longestNiceSubarray(nums []int) int {
	r := 1
	for i := 0; i < len(nums); {
		if len(nums)-i < r {
			break
		}

		nextI := i + 1
		for j := i + 1; j < len(nums); {
			flag := false
			for k := i; k < j; k++ {
				if nums[k]&nums[j] != 0 {
					nextI = k + 1
					flag = true
					if r < j-i {
						r = j - i
					}

					break
				}
			}

			if flag {
				break
			}

			if r < j-i+1 {
				r = j - i + 1
			}

			j++
		}

		i = nextI
	}
	return r
}
