package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {

	// Gidday var, make me an empty p1 of type person thx!
	var p1 person

	// Gday Strings, NewReader the following slice of bytes thx!
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)

	// Heya Jason, NewDecode from/using rdr into the address of p1
	// in its Decode (receiving) manifestation thx!
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
