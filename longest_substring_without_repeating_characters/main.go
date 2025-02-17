package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// Time complexity O(n), space complexity O(n).
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	maxLength, left := 0, 0

	for right := range s {
		index, found := charMap[s[right]]
		if found && index >= left {
			left = index + 1
		}

		charMap[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
