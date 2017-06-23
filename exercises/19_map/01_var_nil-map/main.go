package main

import "fmt"

func main() {

	mygreeting := make(map[string]string)
	mygreeting["breakfast"] = "Hummus on toast"
	mygreeting["lunch"] = "Pan-fried salmon"

	fmt.Println(mygreeting)
}
