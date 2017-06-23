package main

import "fmt"

func main() {
	my_slice := make([]int, 1)
	fmt.Println("here is index0 of slice00 :", my_slice[0])

	my_slice[0] = 7
	fmt.Println("Now we assigned 7 to my_slice[0]", my_slice[0])

	my_slice[0]++
	fmt.Println("Finally, we used my_slice[0]++ to get:", my_slice[0])
}
