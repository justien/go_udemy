package main

import "fmt"

func makeGreeter() func() string {

	x := 4
	fmt.Println("\n")
	fmt.Println("MG: makeGreeter() begins to exist.")
	fmt.Println("MG: I emit the value of a string-returning function.")
	fmt.Println("MG: I have a variable x, with value ", x)
	fmt.Println("\n")

	return func() string {

		n := "Hello world"
		fmt.Println("\n")
		fmt.Println("AF: anonymous function func() begins to exist.")
		fmt.Printf("AF: I emit a %T with value %s", n, n)
		fmt.Println(n)
		fmt.Println("\n")
		return n
	}
}

func main() {

	fmt.Println("\n")
	fmt.Println("MAIN: main() begins to exist.\n")
	greet := makeGreeter()
	fmt.Println("MAIN: I have a variable greet, which is a value with type %T", greet)
	fmt.Println("MAIN: Now we print the value of greet():", greet())
}

/* NOTE WELL:
   as soon as a func expression is declared, any non-returned
   instructions are executed!
*/
