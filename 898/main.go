package main

import "fmt"

func main() {
	fmt.Println(subarrayBitwiseORs([]int{1, 1, 2}))
	fmt.Println(subarrayBitwiseORs([]int{1, 2, 4}))
}

func subarrayBitwiseORs(arr []int) int {
	mp := make(map[int]bool)
	cmp := make(map[int]bool)
	cmp[0] = true
	for i := 0; i < len(arr); i++ {
		nmp := make(map[int]bool)

		nmp[arr[i]] = true
		mp[arr[i]] = true

		for k := range cmp {
			nmp[k|arr[i]] = true
			mp[k|arr[i]] = true
		}
		cmp = nmp
	}

	return len(mp)
}
