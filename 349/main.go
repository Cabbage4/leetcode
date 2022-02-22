package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	mp1 := make(map[int]bool)
	mp2 := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range nums1 {
		mp1[v] = true
	}

	for _, v := range nums2 {
		if mp1[v] && !mp2[v] {
			mp2[v] = true
			result = append(result, v)
		}
	}

	return result
}
