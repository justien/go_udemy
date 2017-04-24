package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like our owl?`
	g := 'M'

	fmt.Printf(" %v is %T \n", a, a)
	fmt.Printf(" \"%v\" is %T \n", b, b)
	fmt.Printf(" %v is %T \n", c, c)
	fmt.Printf(" %v is %T \n", d, d)
	fmt.Printf(" \"%v\" is %T \n", e, e)
	fmt.Printf(" %v is %T \n", f, f)
	fmt.Printf(" %v is %T \n", g, g)
}
