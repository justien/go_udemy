package main

import "fmt"

func main() {
	fmt.Println()
	if false {
		fmt.Println("It was false, so we print the 1st statemnet.")
	} else if true {
		fmt.Println("Else if true - here is the second statement")
	} else {
		fmt.Println("Elsewise - we have made it to 3rd statement")
	}

	fmt.Println()

	if false {
		fmt.Println("If it was false, then be here")
	} else if false {
		fmt.Println("Else if it was false, then go here, where we will never be")
	} else if true {
		fmt.Println("Else if true, please enjoy a candy")
	} else {
		fmt.Println("... at the end end of the if else.  The dust is undisturbed.")
	}
	fmt.Println()
}
