package main

import "fmt"

func reduce(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	fmt.Println(reduce([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}
