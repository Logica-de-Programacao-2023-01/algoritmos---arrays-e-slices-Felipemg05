package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{8, 5, 12, 3, 10, 15, 7, 1, 6, 9}

	min := math.MaxInt64
	max := math.MinInt64

	for _, valor := range slice {
		if valor < min {
			min = valor
		}
		if valor > max {
			max = valor
		}
	}

	fmt.Println("Valor mínimo:", min)
	fmt.Println("Valor máximo:", max)
}
