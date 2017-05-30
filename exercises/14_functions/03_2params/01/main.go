package main

import "fmt"

func main() {
	greet("Darren", "Wolven")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
