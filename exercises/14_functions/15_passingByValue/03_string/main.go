package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println("Name has value of ", name)

	changeMe(name)

	fmt.Println("Name now has value of", name)
}

func changeMe(z string) {
	fmt.Println(z)
	z = "Rocky"
	fmt.Println("z has value", z)
}
