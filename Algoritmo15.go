package main

import (
	"fmt"
)

func main() {
	array := [10]float64{3.2, 8.5, 1.7, 6.9, 4.1, 9.3, 2.8, 7.6, 5.4, 0.9}

	slice := make([]float64, 0)


	for _, valor := range array {
		if valor > 5 {
			slice = append(slice, valor)
		}
	}
	fmt.Println(slice)
}
