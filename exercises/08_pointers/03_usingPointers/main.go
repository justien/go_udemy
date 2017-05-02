package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a is ", a)
	fmt.Println("&a is ", &a)

	var b = &a
	fmt.Println("b is ", b)
	fmt.Println("dereffed b is ", *b)

	*b = 42
	fmt.Println("if dereffed b is set to 42, a is now", a)
}
