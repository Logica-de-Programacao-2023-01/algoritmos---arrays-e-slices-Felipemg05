package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite o tamanho dos arrays: ")
	fmt.Scanln(&n)

	array1 := make([]int, n)
	array2 := make([]int, n)

	fmt.Println("Digite os elementos do primeiro array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scanln(&array1[i])
	}

	fmt.Println("Digite os elementos do segundo array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i)
		fmt.Scanln(&array2[i])
	}

	arraySoma := make([]int, n)
	for i := 0; i < n; i++ {
		arraySoma[i] = array1[i] + array2[i]
	}
	fmt.Println("Array soma:", arraySoma)
}
