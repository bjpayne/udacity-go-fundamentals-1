# UDACITY - Go Fundamentals 1

Variables, conditionals, functions, arrays and slices, loop, and ranges.

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
