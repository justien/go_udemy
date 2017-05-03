package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a is", a)
	fmt.Println("&a is ", &a)

	var b *int = &a
	fmt.Println("b has type *int, is &a, and is", b)

}

/*
b points to the memory address which is used by a
It's not clear to me what b becomes ...
   *int means it's of type 'int pointer'
   - is it now a pointer? or the value at that pointer?
b is of type - 'a pointer to an int'
this means it's a pointer to a memory address where an int is stored
*/
