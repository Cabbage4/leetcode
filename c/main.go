package main

import "fmt"

func main() {
	fmt.Println(goodIndices([]int{123592, 303692, 752301, 522418, 100, 68, 12, 12, 11, 7, 3, 2, 1, 1, 1, 1, 1, 1, 1, 477952, 882691, 945349, 997202, 999245, 999319, 999911, 999913, 999985, 999996, 999996, 999996, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 1000000, 482896, 250330, 20230, 570399, 972746, 999408}, 12))
	//fmt.Println(goodIndices([]int{478184, 863008, 716977, 921182, 182844, 350527, 541165, 881224}, 1))
	//fmt.Println(goodIndices([]int{878724, 201541, 179099, 98437, 35765, 327555, 475851, 598885, 849470, 943442}, 4))
	//fmt.Println(goodIndices([]int{440043, 276285, 336957}, 1))
	//fmt.Println(goodIndices([]int{2, 1, 1, 1, 3, 4, 1}, 2))
}

func goodIndices(nums []int, k int) []int {
	if 2*k == len(nums) {
		return nil
	}

	var r []int
	leftLastIsSort := false
	rightLastIsSort := false
	for i := k; i < len(nums)-k; i++ {
		if i == 42 {
			fmt.Println()
		}
		flag := false

		if !leftLastIsSort {
			for ii := i - 1; ii > i-k; ii-- {
				if nums[ii] > nums[ii-1] {
					flag = true
					break
				}
			}

			if flag {
				rightLastIsSort = false
				continue
			}
			leftLastIsSort = true
		} else {
			if k != 1 && nums[i-1] > nums[i-2] {
				leftLastIsSort = false
				rightLastIsSort = false
				continue
			}
		}

		if !rightLastIsSort {
			for ii := i + 1; ii < i+k; ii++ {
				if nums[ii] > nums[ii+1] {
					flag = true
					break
				}
			}

			if flag {
				continue
			}
			rightLastIsSort = true
		} else {
			if k != 1 && nums[i+k] < nums[i+k-1] {
				rightLastIsSort = false
				continue
			}
		}

		r = append(r, i)
	}

	return r
}
