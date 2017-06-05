package main

import "fmt"

func main() {
	x := 42
	fmt.Println("This is the value of x", x)
	{
		fmt.Println("Indented in one level,, here is x", x)
		y := "Banana cake for breakfast"
		fmt.Println("Here is y's value:", y)
	}
}
