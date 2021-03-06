package main

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	counter := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 4 {
			counter += i
		}
	}
	fmt.Println(counter)
}
