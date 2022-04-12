package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", strStr("akira", "ra")) // expected output: 3
}

// Implement strStr().
// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
// Clarification:
// What should we return when needle is an empty string? This is a great question to ask during an interview.
// For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
func strStr(haystack string, needle string) int {
	length := len(needle)

	if length == 0 {
		return 0
	}

	limit := len(haystack) - length + 1

	for i := 0; i < limit; i++ {
		// search in slices of the same needle length in the haystack
		if haystack[i:i+length] == needle {
			return i
		}
	}

	return -1
}
