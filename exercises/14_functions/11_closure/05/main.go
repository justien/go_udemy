package main

import "fmt"

func wrapper() func() int {
	var x int
	fmt.Println(x)
	return func() int {
		x++
		return x
	}
}

func main() {
	upA := wrapper()
	upB := wrapper()
	fmt.Println("1st increment via upA:", upA())
	fmt.Println("2nd increment via upA:", upA())
	fmt.Println("3rd increment via upA:", upA())
	fmt.Println("Here is upB arriving with a value:", upB())
}
