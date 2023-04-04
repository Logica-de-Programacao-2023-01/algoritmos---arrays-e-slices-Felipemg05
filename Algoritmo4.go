package main

import "fmt"

func main() {
	var slice []int
	var soma, tamanho, x int

	fmt.Print("Digite o tamanho do Slice:")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)
		slice = append(slice, x)
		soma = soma + x
	}

	fmt.Println(slice)
	fmt.Println(soma)

}
