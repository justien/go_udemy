package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Havesham"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

// we use fmt.Sprint() bc greet() returns a single string
// and fmt.Sprint lets us put individual items together
// into a single string.
// It prints items into a single string.
