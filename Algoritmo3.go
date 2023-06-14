package main

import "fmt"

func main() {
    array := [4]float64{2.5, 1.5, 3.0, 4.5}

    produto := 1.0
    for _, valor := range array {
        produto *= valor
    }

    fmt.Println(produto)
}
