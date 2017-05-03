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
	for i < 7 {
		fmt.Println()
		fmt.Printf("i is %d\n", i)
		thisrand := random(i, i*100)
		fmt.Println("this gem shines like the forgotten number ", thisrand)
		i++
	}

}
