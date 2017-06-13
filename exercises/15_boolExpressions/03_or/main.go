package main

import "fmt"

func main() {
	if true || false {
		fmt.Println("This will ALWAYS print")
	}
}
