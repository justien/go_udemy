package main

import "fmt"

func main() {

	myGreet := map[int]string{
		1: "good morning",
		2: "hallo",
		3: "buenos dias",
		4: "buongiorno",
	}

	fmt.Print(myGreet, "\n")

	// delete(myGreet, 2)

	if val, exists := myGreet[2]; exists {
		fmt.Println("val", val, "exists", exists)
	} else {
		fmt.Println("val", val, "exists", exists)
	}

}
