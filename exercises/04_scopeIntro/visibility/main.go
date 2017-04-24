package main

import (
	"fmt"
	"github.com/justien/go_udemy/exercises/04_scopeIntro/visibility/vis"
)

func main() {
	fmt.Println("This is the name found scratched on a rock: ", vis.MyName)
	vis.PrintNames()
	fmt.Println(vis.PrintNum(7))
	fmt.Println("this is the type of func printer: %T ", vis.PrintNames)
	fmt.Println("the type of func PrintNum is %T", vis.PrintNum)
}
