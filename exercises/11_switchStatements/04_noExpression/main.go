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
	fmt.Println()

	gemstash := random(1, 20)
	sisterfriend := "BoneChange"
	switch {
	case gemstash < 10:
		fmt.Println("Looking for gems in your clothing.")
		fmt.Printf("You find %d", gemstash, ".")
	case gemstash >= 10:
		fmt.Println("In your pocket, you find", gemstash, "gems.")
		fmt.Println("Two gems to rub together. Better than yesterday.")
	}
	fmt.Println()

	switch {
	case len(sisterfriend) == 3:
		fmt.Println("The sisterfriend that you use is using you")
	case len(sisterfriend) == 10:
		fmt.Println("The name of the sisterfriend is an approved length")
	case len(sisterfriend) > 7:
		fmt.Println("You struggle to pronounce the name. It is hard to read.")
		fallthrough
	default:
		fmt.Println("The local Office requests information.")
		fmt.Println("You cannot remember where your papers are kept.")

	}
}
