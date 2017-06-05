package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("The 1st increment of x:", increment())
	fmt.Println("The 2nd increment of x:", increment())
}
