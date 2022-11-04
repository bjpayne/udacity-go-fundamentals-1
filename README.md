# UDACITY - Go Fundamentals 1

Variables, conditionals, functions, arrays and slices, loops, and ranges.

### Variables
Typed: `var message string = "Hello World"`

Inferred: `var message = "Hello World"`

Shorthand (inside a functional only): `message := "Hello World`

Example:
```go
package main

import "fmt"

func main() {
	product := "T-shirt"

	cost := 20

	fmt.Println("product's value is:", product)
	fmt.Printf("product's type is: %T\n", product)

	fmt.Println("cost's value is:", cost)
	fmt.Printf("cost's type is: %T\n", cost)
}
```

### Conditionals
```go
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
```

### Functions
```go
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
```

### Arrays and Slices
```go
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
```

### Loops
```go
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
```

### Ranges
```go
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
```