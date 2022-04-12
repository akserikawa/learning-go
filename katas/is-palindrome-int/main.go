package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	print(isPalindrome(1000))
}

func isPalindrome(x int) bool {
	reverse, _ := strconv.Atoi(
		strings.Trim(
			strings.Replace(fmt.Sprint(splitInt(x)), " ", "", -1),
			"[]",
		),
	)
	return x == reverse
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}
