package main

import "fmt"

func main() {
	var x [256]int
	fmt.Println("How long is x?", len(x))

	// let's make a loop 2 to the eight elements long
	// that initialises the empty array
	for i := 0; i < 256; i++ {
		x[i] = i
	}

	// now we loop to make a pattern with the elements of our array
	//... showing all the properties the elements have
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 5 {
			break
		}
	}
}
