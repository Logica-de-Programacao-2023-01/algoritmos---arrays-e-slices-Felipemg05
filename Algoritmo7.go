package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0, 5)

	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	presente := false
	for _, valor := range slice {
		if valor == numero {
			presente = true
			break
		}
	}

	if !presente {
		slice = append(slice, numero)
	}

	fmt.Println(slice)
}
