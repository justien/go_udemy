package main

import "fmt"

// info() is a func that needs a shape
func info(z shape) {
	fmt.Println("The value of me is:", z)
	fmt.Println("The area of me is:", z.area())
}

/*
An interface defines functionality.

shape can interface with methods having the
signature " area() float64 "

Anything that has this method's signature implements the shape interface
*/
type shape interface {
	area() float64
}

/*
area () is a func that needs a square
We attach a method to the user-defined type
 - area() is attached to square

This has the same signature as the shape interface,
therefore it implements the interface
*/
func (z square) area() float64 {
	return z.side * z.side
}

/*
square is a struct that needs a float64
*/
type square struct {
	side float64
}

func main() {

	s := square{10}

	fmt.Printf("The type of variable s is:%T\n", s)
	fmt.Println("We can call info(s):")
	info(s)
	fmt.Println("We can all the method directly against the variable with s.area():", s.area())
}
