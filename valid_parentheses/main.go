package main

import "fmt"

func main() {

	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))

}

func isValid(s string) bool {
	openPairMap := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []int32

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != openPairMap[char] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
