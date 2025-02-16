package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("", "ahbgdc"))
	fmt.Println(isSubsequence("aza", "abzba"))
	fmt.Println(isSubsequence("abbc", "ahbdc"))
}

// Time complexity O(n), Space complexity O(1).
func isSubsequence(s, t string) bool {
	pLeft, bLeft := 0, len(s)
	pRight, bRight := 0, len(t)

	for pLeft < bLeft && pRight < bRight {
		if s[pLeft] == t[pRight] {
			pLeft++
		}

		pRight++
	}

	return pLeft == bLeft
}
