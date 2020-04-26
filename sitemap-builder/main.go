package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url that you want to be a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
}
