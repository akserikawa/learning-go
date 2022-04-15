package main

import "fmt"

func main() {
	fmt.Printf("%v\n", searchInsert([]int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3))
}

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		v := nums[m]
		switch {
		case v < target:
			l = m + 1
		case v > target:
			r = m - 1
		default:
			return m
		}
	}
	return l
}
