package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Hallo Daniel")
	case "Medhi":
		fmt.Println("This is Medhi")
	default:
		fmt.Println("Did you leave?")
	}
}
