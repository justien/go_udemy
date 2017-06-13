package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c0 := customer{"Todd", 44}
	fmt.Println("Here is the memory address of name in c0", &c0.name)

	changeMe(&c0)

	fmt.Println("Here is variable c0", c0)
	fmt.Println("and here is the memory address of name in c0:", &c0.name)
}

func changeMe(z *customer) {
	fmt.Println("Here is z:", z)
	fmt.Println("And here is the memory address of name in z:", &z.name)

	z.name = "rocky"

	fmt.Println("Now here is z:", z)
	fmt.Println("And here is the memory address of name in z:", &z.name)
}
