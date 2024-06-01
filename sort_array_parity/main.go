package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0}))
	fmt.Println()
	fmt.Println(sortArrayByParity2([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity2([]int{0}))
}

func sortArrayByParity2(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 > nums[right]%2 {
			nums[left], nums[right] = nums[right], nums[left]
		}

		if nums[left]%2 == 0 {
			left++
		}

		if nums[right]%2 == 1 {
			right--
		}
	}

	return nums
}

func sortArrayByParity(nums []int) []int {
	even, odd := make([]int, 0, len(nums)), make([]int, 0, len(nums))

	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}

	return append(even, odd...)
}
