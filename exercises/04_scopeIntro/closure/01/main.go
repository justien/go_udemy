package main

import "fmt"

var x int

func increment() int {
	fmt.Println("Here is x at the beginning of increment(): ", x)
	x++
	return x
}

func main() {
	fmt.Println(x)
	fmt.Println("Rain dripped off the leaves", increment())
	fmt.Println("Rain dripped off the leaves", increment())
	fmt.Println("There is no more rain dripping.")
}

/*
|    In this script, x is declared right at the top, so it's
|    able to be addressed by any function in the program.

|    This is called package level scope.
*/
