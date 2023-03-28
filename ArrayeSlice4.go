package main

import "fmt"

func main() {
	numeros := [6]float64{12.64, 17, 32, 4.5, 6.3, 2}
	var soma float64

	for _, x := range numeros {
		soma += x
	}

	media := soma / float64(len(numeros))
	fmt.Println(media)
}
