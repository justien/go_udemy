package main

import "fmt"

func main() {

	running_total := 0

	for i := 0; i < 100; i++ {

		if i%3 == 0 || i%5 == 0 {
			fmt.Println("i is a hit with value", i)
			hit := i
			running_total = running_total + hit
			fmt.Println("Current total of all hits is", running_total)
		}
	}

}
