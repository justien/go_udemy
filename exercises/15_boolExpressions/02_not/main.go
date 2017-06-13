package main

import "fmt"

func main() {

	if !true {
		fmt.Println("You'll never see this")
	}

	if !false {
		fmt.Println("You'll only ever see this")
	}
}
