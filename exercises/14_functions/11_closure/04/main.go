package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println("This is the first call of increment:", increment())
	fmt.Println("And the 2nd call of increment:", increment())
}

/*
... I don't yet understand the point of this wrapper concept.
Here we have an anonymous function inside a named function.
Conceptually, what's the advantage?
The anonymous func's functionality could sit directly inside the
named function.
I suppose this is one way to decompose programmatic actions ...?
It looks modular ... you take all the independant actions and make
each one into its own function ... ?
I'm sure this will make awesome amd fundamental sense later on!
*/
