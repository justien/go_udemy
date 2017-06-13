package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(varying ...int) {

	fmt.Println("foo: foo has come into being")
	fmt.Println("foo: the input was:", varying)
}

/*

Write a function, foo, which can be called in all of these ways:

func main() {
foo(1,2)
foo((1,2,3)
aSlice := []int{1,2,3,4}
foo(aSlice...)
foo()
}


*/
