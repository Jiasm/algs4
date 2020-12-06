package main

import "fmt"

func main() {
	fmt.Println(equal(1.0, 2.0))

	fmt.Println(equal(1.0, 0.5))
}

func equal(a, b float64) bool {
  return a >= 0 && a <= 1 && b >= 0 && b <= 1
}