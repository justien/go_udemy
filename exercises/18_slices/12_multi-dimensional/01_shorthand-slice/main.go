package main

import "fmt"

func main() {

	student := []string{}
	students := [][]string{}

	// student[0] = "Todd"

	student = append(student, "Todd")

	fmt.Println("Here is student", student)
	fmt.Println("Here is students", students)

	fmt.Println("Is the slice student empty?", student == nil)
}
