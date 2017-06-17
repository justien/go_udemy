package main

import "fmt"

func main() {
	var x [256]byte
	fmt.Println("X is this long:", len(x))
	fmt.Println("This is the 0th element of array x:", x[0])

	for i := 0; i < 256; i++ {
		x[i] = byte(i)

	}

	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 5 {
			break
		}
	}

}
