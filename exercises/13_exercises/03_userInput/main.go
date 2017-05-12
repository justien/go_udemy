package main

import "fmt"

func main() {
	var name string

	fmt.Println()
	fmt.Println("How would you like me to name you?")
	fmt.Scan(&name)
	fmt.Printf("Hello %s, welcome to this place.\n", name)
	fmt.Println()
}

/*
This is a truly terrible piece of code to be teaching
to rank beginners without warning about what happens
when the input is more than one word long
*/
