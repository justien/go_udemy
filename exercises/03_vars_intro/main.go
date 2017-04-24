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

	fmt.Printf("Printing variables using the %v (show variable) verb \n")
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	// printing variables using the %T (show type) verb

	fmt.Printf("\n")
	fmt.Printf("Printing variables using the %%T (show type) verb \n")
	fmt.Printf("%v is %T \n", a, a)
	fmt.Printf("\"%v\" is %T \n", b, b)
	fmt.Printf("%v is %T \n", c, c)
	fmt.Printf("%v is %T \n", d, d)
	fmt.Printf("\"%v\" is %T \n", e, e)
	fmt.Printf("%v is %T \n", f, f)
	fmt.Printf("%v is %T \n", g, g)

	// declaring variables with zero value

	var h int
	var i string
	var j float64
	var k bool

	fmt.Printf("\n")
	fmt.Printf("Declaring variables with a zero value \n")
	fmt.Printf("%v \n", h)
	fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", j)
	fmt.Printf("%v \n", k)
}
