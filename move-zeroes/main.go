package main

import "fmt"

func main() {

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	nonZeroIndex := 0

	for _, num := range nums {
		if num != 0 {
			nums[nonZeroIndex] = num
			nonZeroIndex++
		}
	}

	for nonZeroIndex < len(nums) {
		nums[nonZeroIndex] = 0
		nonZeroIndex++
	}
}
