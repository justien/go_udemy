package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Iris":
		fmt.Println("This is Riis")
	case "Token":
		fmt.Println("we eaten lunch")
	case "Marcus":
		fmt.Println("We found Marcus and a fallthrough!")
		fallthrough
	case "A place to rest":
		fmt.Println("We have a place to rest and a fallthrough ...")
		fallthrough
	case "A small pond":
		fmt.Println("You find a frog near the pond")
	default:
		fmt.Println("We reached the default")
	}
}
