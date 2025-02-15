package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge1(nums1, 3, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{1}
	nums2 = []int{}
	merge1(nums1, 1, nums2, 0)
	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	merge1(nums1, 0, nums2, 1)
	fmt.Println(nums1)

	fmt.Println()

	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}
	merge1(nums1, 3, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{1}
	nums2 = []int{}
	merge2(nums1, 1, nums2, 0)
	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	merge2(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}

// Time complexity O(m+n), space complexity O(1).
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}

		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// Time complexity O(m+n), space complexity O(m+n).
func merge1(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0, m+n)

	i, j := 0, 0

	for {
		switch {
		case i < m && j < n:
			if nums1[i] <= nums2[j] {
				result = append(result, nums1[i])
				i++
			} else {
				result = append(result, nums2[j])
				j++
			}
		case i < m:
			result = append(result, nums1[i])
			i++
		case j < n:
			result = append(result, nums2[j])
			j++
		default:
			break
		}
	}

	copy(nums1, result)
}
