package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

/*
We can take a type, and apply a method.
This is like saying: Take p1 and do fullName with it!

In this example, it's possible call fullName without specific variables -
ie p1.fullName() - because fullName() itself has a receiver parameter,
where it states which variables it receives (ie (p person) )
*/
