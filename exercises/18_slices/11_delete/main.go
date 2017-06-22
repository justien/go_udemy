package main

import "fmt"

func main() {

	mySlice00 := []string{"Monday", "Tuesday"}
	myOtherSlice01 := []string{"Wednesday", "Thursday", "Friday"}

	mySlice00 = append(mySlice00, myOtherSlice01...)
	fmt.Println(mySlice00)

	mySlice00 = append(mySlice00[:2], mySlice00[3:]...)
	fmt.Println(mySlice00)

}
