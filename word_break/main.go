package main

import "fmt"

func main() {
	fmt.Println(wordBreak1("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak1("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak1("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak1("aaaaaaa", []string{"aaaa", "aaa"}))
	fmt.Println()
	fmt.Println(wordBreak2("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak2("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak2("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak2("aaaaaaa", []string{"aaaa", "aaa"}))
}

func wordBreak2(s string, wordDict []string) bool {
	dict := make(map[string]struct{})
	for _, str := range wordDict {
		dict[str] = struct{}{}
	}

	queue := []int{0}
	visited := make([]bool, len(s)+1)

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]

		if visited[start] {
			continue
		}
		visited[start] = true

		for end := start + 1; end <= len(s); end++ {
			if _, found := dict[s[start:end]]; found {
				if end == len(s) {
					return true
				}
				queue = append(queue, end)
			}
		}
	}

	return false
}

func wordBreak1(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, str := range wordDict {
		dict[str] = true
	}

	spaces := make([]bool, len(s)+1)
	spaces[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if spaces[j] && dict[s[j:i]] {
				spaces[i] = true
				break
			}
		}
	}

	return spaces[len(s)]
}
