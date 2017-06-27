package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Vorname  string
	Nachname string `json:"-"`            // don't export this KV pair
	Age      int    `json:"wisdom score"` // these tags will replace the keys
}

func main() {

	p1 := person{"James", "Bond", 20}

	myByteSlice, _ := json.Marshal(p1)

	fmt.Println(string(myByteSlice)) //--> {"First":"James","wisdom score":20}
}
