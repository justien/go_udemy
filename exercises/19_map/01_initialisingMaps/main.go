package main

import "fmt"

func main() {

	// make a map with a VAR NIL
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)

	/*
	   	myGreeting["Jenny"] = "Bonjour."
	    and you will get this:
	    panic: assignment to entry in nil map
	*/

	// make a map with a VAR MAKE
	myFood := make(map[string]string)
	myFood["breakfast"] = "Hummus on toast"
	myFood["lunch"] = "Pan-fried salmon"
	fmt.Println(myFood)

	// make a map with SHORTHAND MAKE
	myDrink := make(map[string]string)
	myDrink["am"] = "cup of tea"
	myDrink["day"] = "more cups of tea"
	fmt.Println(myDrink)

	// make a map with COMPOSITE LITERAL
	myNachSpeise := map[string]string{}
	myNachSpeise["Monday"] = "Flat Peach"
	myNachSpeise["Wednesday"] = "Butter Spitzgebäck"
	fmt.Println(myNachSpeise)

	// make a map with INITIALISED COMPOSITE LITERAL
	myNeighbour := map[string]string{
		"4OG": "Friendly single Mum and astronaut daughter",
		"3OG": "Markus and family, nuggety man and family",
		"2OG": "Mannequin A and his airBnB peeps",
	}
	fmt.Println(myNeighbour)
}
