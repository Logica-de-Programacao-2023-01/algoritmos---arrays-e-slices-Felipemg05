package main

import (
	"fmt"
)

func main() {
	array := [5]int{2, 7, 9, 12, 15}

	slice := make([]int, 0)

	for _, valor := range array {
		if valor%3 == 0 {
			slice = append(slice, valor)
		}
	}

	fmt.Println(slice)
}
