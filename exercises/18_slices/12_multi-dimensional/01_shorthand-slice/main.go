package main

import "fmt"

func main() {

	student := []string{}
	students := [][]string{}

	fmt.Println("Here is student", student)
	fmt.Println("Here is students", students)

	fmt.Println("Is the slice student empty?", student == nil)
}
