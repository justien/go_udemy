package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int // all these capitalised field are exported
	notExported int // this is lowercase, and not exported
}

func main() {

	p1 := person{"James", "Bond", 20, 007}
	myByteSlice, _ := json.Marshal(p1)

	fmt.Println(p1)
	fmt.Println(myByteSlice)
	fmt.Printf("%T \n", myByteSlice)
	fmt.Println(string(myByteSlice))
}
