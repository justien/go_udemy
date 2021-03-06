package main

import "fmt"

func main() {

	x := 0

	increment := func() int {
		x++
		return x
	}

	decrement := func() int {
		x := x - 1
		return x
	}

	fmt.Println("The 1st increment of x:", increment())
	fmt.Println("The 2nd increment of x:", increment())

	fmt.Println("The 1st decrement of x:", decrement())
}

/*
closure helps us limit the scope of variables used by multiple functions.
in the above example, x is only available to two functions bc one
has been anonymously embedded inside main()
*/
