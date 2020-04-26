package main

import "fmt"

func main() {
	fmt.Println(countSheep(2)) // 1 sheep...2 sheep...
	fmt.Println(countSheep(3)) // 1 sheep...2 sheep...3 sheep...
}

func countSheep(num int) string {
	var str = ""
	for i := 1; i <= num; i++ {
		str = str + fmt.Sprintf("%d sheep...", i)
	}
	return str
}
