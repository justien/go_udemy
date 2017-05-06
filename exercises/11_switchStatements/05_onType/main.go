package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("The sisterfriend gift is a number.")
		fmt.Println("You think it might be a forgotten one,")
		fmt.Println("left from after the pronunciation of the 99 Names.")
		fmt.Println()
	case string:
		fmt.Println("By the park, a", x, "is held in place by a string")
		fmt.Println()
	case contact:
		fmt.Println("This contains contact information")
	default:
		fmt.Println("Your sisterfriend gave you something.")
		fmt.Println("You cannot recognise what it is.")
		fmt.Println("She hides her love for you by these riddles.")
	}
}

func main() {
	fmt.Println()
	SwitchOnType(7)
	SwitchOnType("Rattling box")
	var t = contact{"A muddied scrap", "from a headdress"}
	SwitchOnType(t)
	fmt.Println(t.greeting)
	SwitchOnType(t.greeting)
	fmt.Println(t.name)
	SwitchOnType(t.name)
	fmt.Println()
}
