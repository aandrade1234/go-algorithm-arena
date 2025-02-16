package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

// Time complexity O(n), space complexity O(1).
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
