package main

import "fmt"

func main() {

	b := true
	c := false
	if food := "Chocolate"; c {
		fmt.Println(food)
	}
	if food := "Kale"; b {
		fmt.Println(food)
	}

	/*
	  Watch out for the false friend - the phrase only mimics an
	  English phrase.
	  The proper meaning is closer to
	  " if food, in having the value "Chocolate", is true, then
	  perform this function.
	*/

	// fmt.Println(food)
}
