package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p1 := &person{"James", 20}

	fmt.Println(p1) // --> Gives &{James 20}

	fmt.Printf("%T\n", p1) // --> Gives *main.person

	fmt.Println(p1.name)
	fmt.Println(p1.age)

}
