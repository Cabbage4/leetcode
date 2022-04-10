package main

func main() {
}

func findFinalValue(nums []int, original int) int {
	mp := make([]bool, 1001)
	for _, v := range nums {
		mp[v] = true
	}
	for {
		if !mp[original] {
			return original
		}

		original *= 2
		if original >= len(mp) {
			return original
		}
	}
}
