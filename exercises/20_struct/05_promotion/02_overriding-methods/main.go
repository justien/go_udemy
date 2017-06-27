package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

/*

There's an important difference between these two functions.
The first one is Greeting() which kicks in when a person is received.
The second one is triggered when a doubleZero looks for a Greeting().

So even though these functions have the same name, there's no
confusion  - their differing receivers distinguish them from
one another.

*/

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {

	p1 := person{"Ian Flemming", 44}

	p2 := doubleZero{
		person:        person{"James Bond", 23},
		LicenseToKill: true,
	}

	/*
	   p1 is of type person.
	     It'll look for functions which receive person types,
	     and will ignore functions that don't receive people.

	   p2 is of type doubleZero.
	     It only pays attention to funcs which receive doubleZero things.
	*/

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
