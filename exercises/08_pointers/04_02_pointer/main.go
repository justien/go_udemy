package main

import "fmt"

func zero(z *int) {
	fmt.Printf("z has value %d and address %p \n", *z, z)
	*z = 0
	fmt.Printf("z is now set to value %d\n", *z)
}

func main() {
	fmt.Println()
	x := 5
	fmt.Printf("x has value %d and address %p \n", x, &x)
	zero(&x)
	fmt.Printf("after zero(&x), x has value %d\n", x)
	fmt.Println()
}
