package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome1("babad"))
	fmt.Println(longestPalindrome1("cbbd"))
	fmt.Println()
	fmt.Println(longestPalindrome2("babad"))
	fmt.Println(longestPalindrome2("cbbd"))
}

// Time complexity O(n^2), Space complexity O(1).
func longestPalindrome1(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	start := 0
	end := 0

	for i := range n {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}

		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}

// Time complexity O(n^2), Space complexity O(n^2).
func longestPalindrome2(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// mark all substrings of length 1 as a palindrome
	for i := range n {
		dp[i][i] = true
	}

	start := 0
	maxLength := 1

	// check for palindromes in substrings larger than 2
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && // border characters are equal
				(length == 2 || dp[i+1][j-1]) {
				dp[i][j] = true

				if length > maxLength {
					start = i
					maxLength = length
				}
			}
		}
	}

	return s[start : start+maxLength]
}
