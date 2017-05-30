package main

import "fmt"

func main() {
	fmt.Println(greet("Alice ", "Aforethought"))
}

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

// the return is a string, AND it's called 's'
// dunno really why you'd want to do this ...
// unless you're returning a lot of things from the one func
// and you want to list them all out ...
