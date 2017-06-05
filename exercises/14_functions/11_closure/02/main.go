package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println("The first increment of x:", increment())
	fmt.Println("The 2nd increment of x:", increment())
}
