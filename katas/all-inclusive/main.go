package main

import (
	"fmt"
)

// Check if all rotations in string are in arr
func main() {
	strng := "bsjq"
	arr := []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}

	fmt.Print(ContainAllRots(strng, arr), "\n")
}

func ContainAllRots(strng string, arr []string) bool {
	counter := 0
	master := []string{}
	for i, _ := range strng {
		master = append(master, strng[i:]+strng[:i])
	}
	for _, j := range master {
		for _, k := range arr {
			if j == k {
				counter++
				break
			}
		}
	}
	return len(strng) == counter
}
