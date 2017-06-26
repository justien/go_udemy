package main

import "fmt"

func main() {

	myGreet := map[int]string{
		1: "good morning",
		2: "hallo",
		3: "buenos dias",
		4: "buongiorno",
	}

	for key, val := range myGreet {
		fmt.Println(key, "-", val)
	}
}
