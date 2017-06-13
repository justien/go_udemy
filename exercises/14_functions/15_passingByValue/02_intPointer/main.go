package main

import "fmt"

func main() {

	age := 44
	fmt.Println("age has the memory location", &age)

	changeMe(&age) //pass a memory address with a stored value

	fmt.Println("age has the memory location", &age)
	fmt.Println("age has the stored value", age)
}

func changeMe(z *int) {

	fmt.Println("\n\n")

	fmt.Println("changeMe() has been called into being")

	fmt.Println("z has the value", z)
	fmt.Println("z has the stored value", *z)

	*z = 24

	fmt.Println("z now has the stored value", *z)
	fmt.Println("z has the memory location", &z)

	fmt.Println("\n\n")
}

/*
PASSING BY VALUE

Once age is declared (a := 44), we now have a memory location with a value.

*/
