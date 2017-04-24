package main

import "fmt"

func main() {
	sum := max(7)
	fmt.Println(sum)
}

func max(x int) int {
	return 42 + x
}
