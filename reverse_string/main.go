package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseInteractive("yoyo master"))
	fmt.Println(reverseRecursive("yoyo master"))
}

func reverseInteractive(str string) string {
	var sb strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteRune(rune(str[i]))
	}

	return sb.String()
}

func reverseRecursive(str string) string {
	if len(str) <= 1 {
		return str
	}

	return reverseRecursive(str[1:]) + string(str[0])
}
