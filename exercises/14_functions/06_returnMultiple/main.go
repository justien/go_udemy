package main

import "fmt"

func main() {
	fmt.Println(greet("Nnedi", "Okorafur", " "))
}

func greet(fname, lname, seperator string) (string, string) {
	return fmt.Sprint(fname, seperator, lname), fmt.Sprint(lname, seperator, fname)
}
