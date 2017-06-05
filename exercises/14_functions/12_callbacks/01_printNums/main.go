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




*/
