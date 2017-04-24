package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Even gaff tape couldn't keep out the desert grit"
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y
}
