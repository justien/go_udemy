package main

import "fmt"

func main() {

	if true {
		fmt.Println("this is true")
	}

	if false {
		fmt.Println("this is false")
	}

	if !true {
		fmt.Println("this is true")
	}

	if !false {
		fmt.Println("this is false")
	}
}
