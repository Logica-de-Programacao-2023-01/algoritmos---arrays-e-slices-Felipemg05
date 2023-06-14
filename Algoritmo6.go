package main

import (
	"fmt"
)

func main() {
	array := [10]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}

	var valor int
	fmt.Print("Digite um valor: ")
	fmt.Scanln(&valor)

	encontrado := false
	for _, num := range array {
		if num == valor {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("O valor", valor, "está presente no array.")
	} else {
		fmt.Println("O valor", valor, "não está presente no array.")
	}
}
