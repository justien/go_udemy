package main

import "fmt"

func main() {

	theBig := whoBig(2, 5, 3, 7, 9, 45, 23, 47, 0)

	fmt.Println("the Biggest number is", theBig)
}

func whoBig(sf ...int) int {

	meBig := 0

	for _, i := range sf {
		if i > meBig {
			meBig = i
		}
	}

	return meBig
}

/*
Write a function with one variadic parameter that finds the greatest number in a list of numbers.
*/
