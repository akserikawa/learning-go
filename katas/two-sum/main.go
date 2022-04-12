package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if v, found := m[target-num]; found {
			print(m[target-num], "\n")
			return []int{v, idx}
		}
		m[num] = idx
	}
	return nil
}
