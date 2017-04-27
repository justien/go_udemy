package main

import "fmt"

const c = 42

func main() {
	a := 43
	b := c
	d := c
	e := c
	f := c
	g := c

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("a's memory address is ", &a)
	fmt.Printf("%d \n", &a)
	fmt.Println("b's memory address is ", &b)
	fmt.Printf("%d \n", &b)
	fmt.Println("d's memory address is ", &d)
	fmt.Printf("%d \n", &d)
	fmt.Println("e's memory address is ", &e)
	fmt.Printf("%d \n", &e)
	fmt.Println("f's memory address is ", &f)
	fmt.Printf("%d \n", &f)
	fmt.Println("g's memory address is ", &g)
	fmt.Printf("%d \n", &g)
}
