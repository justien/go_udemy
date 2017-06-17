package main

import "fmt"

func main() {
	var x [256]string
	fmt.Println("X is this long", len(x))
	fmt.Println("This is the zeroth element of x:", x[0])

	for i := 0; i < 256; i++ {
		x[i] = string(i)
	}

	for _, v := range x {
		fmt.Printf("%v - %t - %v\n", v, v, []byte(v))
	}
}
