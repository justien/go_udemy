package main

import "fmt"

func main() {

	var student []string
	var students [][]string

	// student[0] = "Todd"

	student = append(student, "Todd")

	fmt.Println("Here is students", students)
	fmt.Println("Here is student", student)

	fmt.Println("Is the slice student empty?", student == nil)
}
