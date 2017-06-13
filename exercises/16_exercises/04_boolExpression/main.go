package main

import "fmt"

func main() {

	x := true && true
	y := false && true
	z := !(false && false)

	ä := x || y || z

	fmt.Println("true && false) || (false && true) || !(false && false) evaluates to", ä)

}

/*

What's the value of this expression: (true && false) || (false && true) || !(false && false)?

*/
