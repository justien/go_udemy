package main

import "fmt"

func ordinals(x int) string {
	x = x % 10
	if x == 1 {
		return "st"
	} else if x == 2 {
		return "nd"
	} else if x == 3 {
		return "rd"
	} else {
		return "th"
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		suffix := ordinals(i)
		if i%10 == 0 {
			fmt.Println()
		}
		if i%2 == 0 {
			fmt.Println("This number is the ", i, suffix)
		}
	}

}
