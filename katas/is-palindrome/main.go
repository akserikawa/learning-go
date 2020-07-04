package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsPalindrome("radar"))
}

func IsPalindrome(str string) bool {
	return strings.EqualFold(str, reverse(str))
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
