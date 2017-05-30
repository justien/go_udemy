package main

import "fmt"

func main() {
	greet("Darren", "Wolven")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}

// here we've shrunk down the parameter declaration list
// by comma-seperating the variable names
