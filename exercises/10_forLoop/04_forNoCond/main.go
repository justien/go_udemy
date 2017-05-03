package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min

}

func main() {
	i := 1
	for i < 3 {
		i++
		if i == 1 {
			fmt.Println("\n\nIt is the first day")
		} else if i == 2 {
			fmt.Println("\n\nIt is the second day")
		} else {
			fmt.Println("\n\n... many, many days")
		}

		fmt.Println("walls pitted from many years of grit-bearing winds")
		fmt.Println("you pass by each day, sometimes a group of children play there")
		fmt.Println("you see their shadows, when the sun shines")
		fmt.Println("but often the shadows are there, alone, without sun or children")
		fmt.Println()
		fmt.Println()
	}

	shadowrand := random(1, 1000)
	fmt.Printf("After exposure to %d shadows, your eyes leak glittering slime.\n\n", shadowrand)
}
