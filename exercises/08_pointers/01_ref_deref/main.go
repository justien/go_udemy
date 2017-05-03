package main

import "fmt"

func main() {
	a := 43

	fmt.Printf("a \tis a variable with value of ")
	fmt.Println("\t", a)                        // print out a
	fmt.Println("the memory address &a is", &a) // a's memory address

	var b *int = &a // b is a pointer to the address of a
	fmt.Printf("b is of type *int and has value ")
	fmt.Println(b) // print out b
	fmt.Printf("*b is the value at memory location b and is ")
	fmt.Println(*b) // print out the value at b
	fmt.Printf("&b is the memory location of b and is")
	fmt.Println(&b)

	//var c *int = &b
	//fmt.Println("c is given the memory address of b ", c)

}

/*
b points to the memory address which is used by a
It's not clear to me what b becomes ...
   *int means it's of type 'int pointer'
   - is it now a pointer? or the value at that pointer?
b is of type - 'a pointer to an int'
this means it's a pointer to a memory address where an int is stored

... * is an operator, like - or + or <
*/
