package main

import (
	"fmt"
	"strings"
)

func add(n1, n2 int) int {
	return n1 + n2
}

func sayHello(name string) string {
	return "Hello" + name
}

func sayLoudly(phrase string, times int) string {
	saidLoudly := ""

	for i := 0; i < times; i++ {
		saidLoudly += strings.ToUpper(phrase) + "!"
	}

	return saidLoudly
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sayHello("Ben"))
	fmt.Println(sayLoudly("Hello", 3))
}
