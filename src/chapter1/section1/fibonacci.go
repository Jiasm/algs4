package main

import "fmt"

func main() {
	f := 0
	g := 1

	for i :=0; i <= 15; i++ {
		fmt.Println(f)

		f, g = f + g, f
	}
}