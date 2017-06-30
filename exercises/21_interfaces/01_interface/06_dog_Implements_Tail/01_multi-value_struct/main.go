package main

import (
	"fmt"
	"math"
)

// Structs -------------------------------------
type square struct {
	side   float64
	colour string
}

type circle struct {
	radius  float64
	emotion string
}

// Interfaces --------------
type dimension interface {
	area() float64
}

type shadow interface {
	tell_me() string
}

// Methods ------------------

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// func (s square) tell_me() string {
// 	w := s.colour
// 	return w
// }

func (s circle) tell_me() string {
	w := fmt.Sprintf("%v", s.emotion)
	return w
}

type myTest interface {
	tell_me() string
}

// functions ------------------------------------
func info(a dimension) {
	fmt.Println("The area is:", a.area())
}

// I don't get why I can't access circle.emotion's value

func the_cloud(a shadow) {
	c := a.tell_me
	fmt.Println("Here is c", c)
	fmt.Printf("Here is value of c %v", c)

	//fmt.Println(fmt.Sprintf("You step out into the shadow of a %v cloud \n", a.tell_me())
	fmt.Println()
}

// the headlining act ---------------------------
func main() {
	s := square{10, "red"}
	c := circle{5, "recriminative"}
	info(s)
	the_cloud(c)
	/* --> the concrete type implementing the shape interface is of type circle
	 */
}
