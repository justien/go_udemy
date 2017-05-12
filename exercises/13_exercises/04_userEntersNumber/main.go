package main

import "fmt"

func main() {

	var theNumber int
	var largerNumber int

	fmt.Println()
	fmt.Println("I would like you to tell me something.")
	fmt.Println("Tell me, today, tell me a number ...")
	fmt.Scan(&theNumber)
	fmt.Println("Thank-you! \nCould you tell me another number, one that's larger?")
	fmt.Scan(&largerNumber)
	remainder := largerNumber % theNumber
	fmt.Println("If I divide the larger by the smaller, what remains is", remainder)
	fmt.Println()
}
