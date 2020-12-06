package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cur int

	first, _ := strconv.Atoi(os.Args[1])

	for _, v := range os.Args[2:4] {
		cur, _ = strconv.Atoi(v)
		
		if cur != first {
			fmt.Println("not equal")
			return
		} 
	}

	fmt.Println("equal")
}