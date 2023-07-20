package main

import (
	"fmt"
	"regexp"
	"strings"
)

// https://www.coderbyte.com/editor/Longest%20Word:Go
//
// Have the function LongestWord(sen) take the sen parameter being passed and return the longest word in the string.
// If there are two or more words that are the same length, return the first word from the string with that length.
// Ignore punctuation and assume sen will not be empty. Words may also contain numbers, for example "Hello world123 567"
//
// Examples
// Input: "fun&!! time"
// Output: time
//
// Input: "I love dogs"
// Output: love

func main() {
	input := "I love dogs"

	output := solution1(input)
	fmt.Println(output)
}

// solution1 time complexity O(n) and space complexity is O(m) where m is the number of words created by split
func solution1(sen string) string {
	output := ""
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	words := strings.Split(sen, " ")

	for _, word := range words {
		cleanWord := reg.ReplaceAllString(word, "")
		if len(cleanWord) > len(output) {
			output = cleanWord
		}
	}

	return output
}
