package main

import "fmt"

func main() {
	switch "frog" {
	case "pond", "tree":
		fmt.Println("A tree overshadows the pond")
	case "dusk", "frog":
		fmt.Println("The pond moistens the frog's skin")
	}
}
