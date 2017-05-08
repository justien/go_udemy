package main

import "fmt"

func main() {

	// b := true
	c := false
	d := 17
	e := 17

	if food := "Chocolate"; c {
		fmt.Println(food)
	}
	if numbers := "twins since birth"; d == e {
		fmt.Println(numbers)
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
