package main

import "fmt"

func main() {
	language := "Go"

	if language == "Java" {
		fmt.Println("language is Java")
	} else if language == "C" {
		fmt.Println("language is Go")
	} else {
		fmt.Println("language not found")
	}
}
