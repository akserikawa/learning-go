package main

import (
	"fmt"
	"strconv"
	"strings"
)

type category struct {
	Letter   string
	Quantity int
}

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}

	expected := "(A : 0) - (B : 1290) - (C : 515) - (D : 600)"

	fmt.Println(expected == StockList(b, c))
}

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	stock := make(map[string]int)
	for _, artCode := range listArt {
		code := strings.Split(artCode, " ")
		letter := strings.Split(code[0], "")[0]
		quantity, _ := strconv.Atoi(code[1])
		stock[letter] += quantity
	}

	result := make([]string, len(listCat))
	for i, letter := range listCat {
		result[i] = fmt.Sprintf("(%s : %v)", letter, stock[letter])
	}

	return strings.Join(result, " - ")
}
