package main

import "fmt"

func main() {
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
	fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1))
}

func LongestConsec(strarr []string, k int) string {
	maxLen := -1
	var longest string
	var longestPos int
	for pos, v := range strarr {
		if len(v) > maxLen {
			longest = v
			maxLen = len(v)
			longestPos = pos
		}
	}
	for i := longestPos + 1; i <= k; i++ {
		longest += strarr[i]
	}
	return longest
}
