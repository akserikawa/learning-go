package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GetCount("abecii"))
}

func GetCount(str string) (count int) {
	s := strings.Split(str, "")
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, v := range s {
		if stringInSlice(v, vowels) {
			count++
		}
	}
	return
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
