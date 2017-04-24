package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("The rain dripped from the gutters ", increment())
	fmt.Println("The rain dripped from the gutters ", increment())
	fmt.Println("There is no more rain dripping.")

}

/*
|    In this script, x is declared and initialised at the start
|    of func main().  A func defined within func main() is
|    also able to use variable x

|    x is no longer available at the package level scope.
|    It's good, because we've narrowed the scope of x.
|    But it can make our code a bit more complex.

|    Also! on line 8 we have an ANONYMOUS function
|    to which we assign the name increment.

|    According to Todd, this is what u do when you need to
|    put a function inside a function.
*/
