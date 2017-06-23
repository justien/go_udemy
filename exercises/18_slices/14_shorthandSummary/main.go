package main

import "fmt"

func main() {

	var student []string

	fmt.Println("Using var student []string, student has the value:", student)
	fmt.Println()

	assistant := []string{}
	fmt.Println("Using assistant := []string{}, assistant has the value:", assistant)
	fmt.Println()

	friend := make([]string, 150, 1000)
	fmt.Println("Using friend := make([]string, 150, 1000) friend has the value:", friend)

}
