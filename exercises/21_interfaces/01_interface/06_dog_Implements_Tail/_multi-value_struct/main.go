package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
	//colour string
}

type circle struct {
	radius  float64
	emotion string
}

// Dimension interface and methods --------------
type dimension interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//shadow interface and methods ------------------
type shadow interface {
	tell_me() string
}

// func (s square) tell_me() string {
// 	w := s.colour
// 	return w
// }

type myTest interface {
	tell_me() string
}

func (s circle) tell_me() string {
	w := s.emotion
	return w
}

// functions ------------------------------------
func info(a dimension) {
	fmt.Println(a)
}

func the_cloud(a shadow) {
	fmt.Println("You step out into the shadow of a", a, "cloud")
}

func main() {
	s := square{10}
	c := circle{5, "recriminative"}
	info(s)
	the_cloud(c)
	/* --> the concrete type implementing the shape interface is of type circle
	 */
}
