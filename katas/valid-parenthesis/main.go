package main

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

func main() {
	fmt.Printf("%v\n", isValid("(){}[]"))
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for _, c := range s {
		if char, ok := m[c]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0
}
