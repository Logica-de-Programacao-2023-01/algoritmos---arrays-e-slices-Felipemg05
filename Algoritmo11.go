package main

import (
	"fmt"
)

func main() {
	array := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var linha, coluna int
	fmt.Print("Digite o índice de linha: ")
	fmt.Scanln(&linha)
	fmt.Print("Digite o índice de coluna: ")
	fmt.Scanln(&coluna)

	if linha >= 0 && linha < len(array) && coluna >= 0 && coluna < len(array[linha]) {
		valor := array[linha][coluna]
		fmt.Println("O valor armazenado nessa posição é:", valor)
	} else {
		fmt.Println("Os índices informados são inválidos.")
	}
}
