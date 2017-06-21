package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 5, 8}

	for i, value := range mySlice {
		fmt.Println(i, " -", value)
	}
}
