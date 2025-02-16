package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var regex = regexp.MustCompile("[^a-z0-9]+")

func main() {
	fmt.Println(isPalindrome1("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome1("race a car"))
	fmt.Println(isPalindrome1("aa"))
	fmt.Println(isPalindrome1(" "))
	fmt.Println()
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome2("race a car"))
	fmt.Println(isPalindrome2("aa"))
	fmt.Println(isPalindrome2(" "))
}

// Time complexity O(n), space complexity O(1)
func isPalindrome2(s string) bool {
	s = strings.ToLower(s)

	left := 0
	right := len(s) - 1

	for left < right {
		rLeft := rune(s[left])
		rRight := rune(s[right])

		if !isValidRune(rLeft) {
			left++
			continue
		}

		if !isValidRune(rRight) {
			right--
			continue
		}

		if rLeft != rRight {
			return false
		}

		left++
		right--
	}

	return true
}

func isValidRune(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

// Time complexity O(n), space complexity O(1)
func isPalindrome1(s string) bool {
	s = strings.ToLower(s)
	s = regex.ReplaceAllString(s, "")

	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
