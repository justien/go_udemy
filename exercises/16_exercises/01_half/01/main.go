package main

import "fmt"

func main() {

	// bingo := half(4)

	fmt.Println(half(7))
}

func half(z int) (int, bool) {
	fmt.Println("Z is", z)
	divide := z / 2
	fmt.Println("z divided by 2 is ", divide)
	remainder := z % 2
	fmt.Println("z remainder 2 is", remainder)
	even := true
	if remainder == 1 {
		even = false
	}
	return divide, even
}

/*
Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2. The second return should be a bool that let’s us know whether or not the argument was even. For example:
half(1) returns (0, false)
half(2) returns (1, true)
*/
