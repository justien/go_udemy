package main

import "fmt"

func main() {
	greet("Horace", "Humbert")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
