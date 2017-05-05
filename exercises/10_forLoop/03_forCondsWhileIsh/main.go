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

func Ordinals(i int) string {

	if i == 1 {
		return "st"
	} else if i == 2 {
		return "nd"
	} else if i == 3 {
		return "rd"
	} else {
		return "th"
	}
}

func main() {
	i := 1
	for i < 7 {
		fmt.Println()
		ordinalsuffix := Ordinals(i)
		fmt.Printf("the %d%s", i, ordinalsuffix)
		fmt.Printf(" bone shines like the abandoned number ")
		thisrandmin := random(i, i*100)
		// fmt.Println(thisrandmin)
		thisrandmax := random(i*100, i*1000)
		// fmt.Println(thisrandmax)
		thisrand := random(thisrandmin, thisrandmax)
		fmt.Println(thisrand)
		i++
	}

}
