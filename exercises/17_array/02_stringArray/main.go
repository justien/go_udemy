package main

import "fmt"

func main() {

	// this collection has 58 items - it is an ARRAY !!!
	var x [58]string

	fmt.Println("Here is x:", x)
	fmt.Println("And this is the length of x:", len(x))
	fmt.Println("Let me show you the 42nd element of array x:", x[42])

	// loop from 65 to 122, initialising the array with values
	// bc the array is type string, these numbers will be ASCII codes
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println("Here is x:", x)
	fmt.Println("And this is the length of x:", len(x))
	fmt.Println("Let me show you the 42nd element of array x:", x[42])

}
