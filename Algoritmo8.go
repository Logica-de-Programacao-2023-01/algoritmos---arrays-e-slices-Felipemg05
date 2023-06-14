package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 8)

	var valor string
	fmt.Print("Digite um valor: ")
	fmt.Scanln(&valor)

	slice[0] = "banana"
	slice[1] = "maçã"
	slice[2] = "banana"
	slice[3] = "laranja"
	slice[4] = "uva"
	slice[5] = "banana"
	slice[6] = "abacaxi"
	slice[7] = "banana"

	novoSlice := make([]string, 0, len(slice))
	for _, elemento := range slice {
		if elemento != valor {
			novoSlice = append(novoSlice, elemento)
		}
	}

	fmt.Println(novoSlice)
}
