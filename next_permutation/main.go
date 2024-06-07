package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums[i+1:])
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
