package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person         // initialises the map with empty strings
	fmt.Println(p1.First) // -->
	fmt.Println(p1.Last)  //-->
	fmt.Println(p1.Age)   //--> 0

	// Cut me a slice of bytes as per this list, thx!
	// enclose in BACKTICKS coz inside I got some doublequotes
	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)

	//Hey Jason, Unmarshall this byteslice to element p1 thx!
	// oh yeah m8 here's the address of p1 btw thx!
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1.First)   //--> James
	fmt.Println(p1.Last)    //--> Bond
	fmt.Println(p1.Age)     //--> 20
	fmt.Printf("%T \n", p1) //--> main.person
}
