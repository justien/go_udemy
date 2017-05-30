package main

import "fmt"

func main() {
	greet("Jane")
	greet("Shayla")
}

// name is only scoped within func greet
func greet(name string) {
	fmt.Println(name)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
