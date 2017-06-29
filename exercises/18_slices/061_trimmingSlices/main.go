package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("%q", strings.Trim(" !!! gloop! gloop! !!! ", "! "))

	a := []string{"a", "b"}
	fmt.Print(a)
	fmt.Print(strings.Trim(fmt.Sprint(a), "[]"))

	//fmt.Printf("%q", strings.Trim(a, "[]"))
	//fmt.Printf(a)
}

/*
Sources:
search:Print Slice without delimiters
https://play.golang.org/p/sk6JKiN6qQ
