package main

import "fmt"

func main() {

	fmt.Println()
	n := average(0, 3, 17, 19, 23, 47, 3.4)
	if n == 0 {
		fmt.Println("No numbers to average")
	} else {
		fmt.Println("The average of the numbers is ", n)
	}
	fmt.Println()
}

func average(sf ...float64) float64 {

	fmt.Println("the values to be averaged are: ", sf)
	fmt.Printf("the type of the values are %T\n", sf)

	//test for zero input
	if len(sf) == 0 {
		return 0
	} else {
		total := 0.0
		// var total float64
		for _, v := range sf {
			total += v
		}
		return total / float64(len(sf))
	}
}

/*
VARIADIC FUNCTION PARAMTERS
in this notation, ...float64 as a parameter prefix means the
function will accept zero or more arguments of type float64.
*/
