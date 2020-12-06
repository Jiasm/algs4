package main

import "fmt"

func main() {
  arr := [][]bool{
		{ true, true, true, false },
		{ true, true, false, true },
		{ true, false, true, true },
		{ false, true, true, true },
	}

	print(arr)
}

func print(arr [][]bool) {
	fmt.Printf("  ")
	for colId := range arr[0] {
		fmt.Printf("%d", colId + 1)
	}

	fmt.Println()

	for rowId, row := range arr {
		fmt.Printf("%d:", rowId + 1)

		for _, value := range row {
			if value {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}

		fmt.Println()
	}
}