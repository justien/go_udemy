package main

import "fmt"

func main() {
	if true && false {
		fmt.Println("This will NEVER print. Ever. Never")
	}
}
