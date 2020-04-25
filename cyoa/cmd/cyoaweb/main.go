package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/akserikawa/learning-go/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
