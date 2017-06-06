package main

import "fmt"

func visit(numbers []int, callback func(int)) {

	fmt.Println("VISIT: visit() has come into being")
	fmt.Println("VISIT: arguments are:", numbers, callback)

	for _, n := range numbers {
		fmt.Println("VISIT_loop: n is value:", n)
		callback(n)
	}

}

func main() {

	fmt.Println("MAIN: main() has come into being")

	visit([]int{1, 2, 3, 5}, func(r int) {

		fmt.Println("VISIT_argument: this is the value of r:", r)

	})
}

/*
when visit() is called, it is passed two arguments:

1. a slice of ints {1, 2, 3, 5}
2. a little function which uses an int called r
   func(r int) {
     fmt.Println("Hello! r has value", r)
   }

... notice that this is an anonymous function with one
parameter - an int called r - and all it does is print the
argument of that parameter.

... so when callback(n) happens in visit(), it performs the function using the current value of n.


*/
