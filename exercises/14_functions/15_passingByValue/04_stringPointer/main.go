package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println("Variable name ha memory addres", &name)

	changeMe(&name)

	fmt.Println("Variable name now has stored value", name)
}

func changeMe(z *string) {
	fmt.Println("z has value of", z)
	*z = "Joshua"
}
