package main

import "fmt"

func hello(num int) {
	fmt.Println("hello", num, " ")
}

func world(num int) {
	fmt.Println("world", num, " ")
}

func main() {
	fmt.Print("line 00:")
	hello(0)
	defer world(0)
	hello(1)
	defer world(1)
	hello(2)
	hello(3)
	hello(4)
	defer world(2)
	defer fmt.Println("The first deferred println")
	defer fmt.Println("The second deferred println")
	defer world(3)
	fmt.Println("The final line. Now to wind out the stack")
}

/*
DEFER

deferred functions
... will wait until the completion of all the other
functions in the block before executing.

they are place on the stack in order of program execution,
and following the final non-deferred instruction, are then
executed in LIFO order.
*/
