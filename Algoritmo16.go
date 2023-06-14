package main

import (
	"fmt"
)

func main() {
	array := [10]int{3, 8, 1, 6, 4, 9, 2, 7, 5, 12}

	slice := make([]int, 0)

	for _, valor := range array {
		if valor%2 == 0 {
			slice = append(slice, valor)
		}
	}

	fmt.Println(slice)
}
