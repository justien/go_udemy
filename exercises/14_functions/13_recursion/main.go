package main

import "fmt"

func factorial(x int) int {
	fmt.Println("Factorial is called into being")
	fmt.Println("x has value:", x)
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}
