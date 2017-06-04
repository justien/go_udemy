package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("helloo world")
	}

	fmt.Println()
	greeting()
	fmt.Printf("%T\t", greeting)
	fmt.Println()
	fmt.Println()
}
