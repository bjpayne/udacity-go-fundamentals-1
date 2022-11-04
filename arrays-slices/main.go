package main

import (
	"fmt"
)

func main() {
	fib := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	subFib := append(fib[0:], fib[len(fib)-2]+fib[len(fib)-1])

	fmt.Println(fib)
	fmt.Println(fib[0])
	fmt.Println(fib[len(fib)-1])
	fmt.Println(len(fib))
	fmt.Println(subFib)
}
