package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Printf("x is 13 %% 3 and has value %d\n", x)
	fmt.Printf("This value is ")
	if x == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}

	for i := 1; i < 7; i++ {
		if i%2 == 1 {
			fmt.Println("i has an odd value")
		} else {
			fmt.Println("i has an even value")
		}
	}
}
