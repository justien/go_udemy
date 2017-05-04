package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		for i := 1; i <= 3; i++ {
			fmt.Println("  ", i)
		}
	}
}
