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

func wrapper(n int) func() int {
	x := 0
	return func() int {
		x++
		gemsum := x + n
		return gemsum
	}
}

func main() {
	wastegems := random(2, 6)
	moregems := wrapper(wastegems)

	fmt.Println()
	fmt.Println()
	fmt.Printf("Whilst crossing the Wastes, you gathered %d gems.\n", wastegems)
	fmt.Printf("You threw one in the fire; it split and became two.\n")
	fmt.Printf("Raking them from the ashes, you then had %d gems.\n", moregems())
	fmt.Printf("You threw another on the coals, risking the attention of minor demons.\n")
	fmt.Printf("Luck held true; you counted the gems and had %d.\n", moregems())
	fmt.Println()
	fmt.Println()

	sibgems := random(9, 37)
	evenmoregems := wrapper(sibgems)
	fmt.Printf("In another hive, a sisterfriend inherited %d gems.\n", sibgems)
	fmt.Printf("By her bed, she found a gem left by a generous lover; now she has %d gems.\n", evenmoregems())
	fmt.Println()
	fmt.Println()
}
