package algorithms

import (
	"regexp"
	"strings"

	"github.com/aandrade1234/go-algorithm-arena/challenges"
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

type LongestWord struct {
	challenges.BaseChallenge
}

func (*LongestWord) GetName() string {
	return "longest-word(source: coderbyte)"
}

func (b *LongestWord) Execute() error {
	return b.BaseExecute("solution 1", b.GetInputs(), b.GetOutputs(), b.solution1)
}

// solution1 time complexity O(n) and space complexity is O(m) where m is the number of words created by split
func (*LongestWord) solution1(input any) any {
	sen := input.(string)

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

func (*LongestWord) GetInputs() []any {
	return []any{
		"fun&!! time",
		"I love dogs",
	}
}

func (*LongestWord) GetOutputs() []any {
	return []any{
		"time",
		"love",
	}
}
