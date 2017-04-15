package main

import "fmt"

// func main() {
//	fmt.Printf("%d \t %b \t %x \n", 42, 42, 42)
//	fmt.Printf("%d \v %b \v %X \n", 42, 42, 42)
// }

func main() {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %q \n", i, i, i)
	}
}
