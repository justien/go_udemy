package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)

	student[0] = "Todd"

	// student = append(student, "Todd")

	fmt.Println("Here is students", students)
	fmt.Println("Here is student", student)

	fmt.Println("Is the slice student empty?", student == nil)
}
