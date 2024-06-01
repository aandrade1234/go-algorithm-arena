package main

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		switch {
		case sum == target:
			return []int{left + 1, right + 1}
		case sum < target:
			left++
		default:
			right--
		}
	}

	return []int{}
}
