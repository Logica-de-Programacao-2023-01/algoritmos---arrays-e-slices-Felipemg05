package main

import (
	"fmt"
)

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	soma := 0

	for i, valor := range array {
		if i%2 == 0 {
			soma += valor
		}
	}

	fmt.Println("A soma dos elementos nas posições pares é:", soma)
}
