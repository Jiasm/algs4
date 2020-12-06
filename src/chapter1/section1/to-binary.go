package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(toBinary(5, ""))
}

func toBinary(num int, str string) string {
	if (num <= 0) {
		return str
	}

	return toBinary(num / 2, strconv.Itoa(num % 2) + str)
}