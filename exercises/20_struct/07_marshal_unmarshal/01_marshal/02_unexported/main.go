package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first string
	last  string
	age   int // none of these are uppercase, so none are exported
}

func main() {

	p1 := person{"James", "Bond", 20}

	fmt.Println(p1)

	bs, _ := json.Marshal(p1)

	fmt.Println(string(bs)) // --> {} because no values could be exported
}
