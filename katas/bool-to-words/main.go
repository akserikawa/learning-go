package main

import (
	"fmt"
)

func main() {
	fmt.Println(BoolToWord(true))  // "Yes"
	fmt.Println(BoolToWord(false)) // "No"
}

func BoolToWord(word bool) string {
	return map[bool]string{false: "No", true: "Yes"}[word]
}
