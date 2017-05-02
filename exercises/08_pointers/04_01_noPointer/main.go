package main

import "fmt"

func zero(z int) {
	fmt.Printf("Pointer address of z in zero() is %p\n", &z)
	fmt.Println("address of z in zero() is", &z)
	z = 0
}

func main() {
	fmt.Println()
	x := 5
	fmt.Printf("Pointer address of x in main() is %p\n", &x)
	zero(x)
	fmt.Println("x in main()  has value", x)
	fmt.Println()
}
