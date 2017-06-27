package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("Here is the type of foo: %T", myAge, "\nand the value of foo: %v", myAge, "\n")
}

/*
Normally, you'd never alias a type this way because it's
redundant.
However if you ever need to attach a method to a type,
then the type must be aliased.
*/
