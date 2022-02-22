package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		value := numbers[left] + numbers[right]
		if value == target {
			return []int{left, right}
		} else if value > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
