package main

import "fmt"

func main() {
	a := 43

	fmt.Println(" a is ", a)
	fmt.Println("&a is ", &a)

	var b *int = &a
	fmt.Println("  b is *int of &a")
	fmt.Println(" b is ", b)
	fmt.Println("*b is ", *b)
	fmt.Println("&b is ", &b)

	*b = 42
	fmt.Println("*b is now ", *b)
	fmt.Println(" a is now ", a)

	a = 41
	fmt.Println(" a is now ", a)
	fmt.Println("*b is now ", *b)
}

/*
42 = *b =
*/
