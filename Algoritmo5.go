package main

import "fmt"

func main() {
	var matriz [3][2]int

	for linha := 0; linha < len(matriz); linha++ {
		for coluna := 0; coluna < len(matriz[linha]); coluna++ {
			var x int
			fmt.Print("dIGITE O NÚMERO DA LINHA", linha, "e da coluna ", coluna)
			fmt.Scan(&x)
			matriz[linha][coluna] = x
		}
	}
	fmt.Println(matriz)
}
