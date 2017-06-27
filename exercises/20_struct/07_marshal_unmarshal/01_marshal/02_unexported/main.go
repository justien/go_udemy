package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first string
	last  string
	age   int // none of these are uppercase, so none will be exported
}

func main() {

	p1 := person{"James", "Bond", 20}

	fmt.Println(p1) // --> will display {James Bond 20}

	bs, _ := json.Marshal(p1)

	fmt.Println(string(bs)) // --> {} because no values could be exported
}
