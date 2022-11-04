package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	output := []string{}

	i := 1

	for ; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			output = append(output, "FizzBuzz")

			continue
		}

		if i%3 == 0 {
			output = append(output, "Fizz")

			continue
		}

		if i%5 == 0 {
			output = append(output, "Buzz")

			continue
		}

		output = append(output, strconv.Itoa(i))
	}

	return output
}

func main() {
	output := fizzBuzz(15)

	fmt.Println(output)
}
