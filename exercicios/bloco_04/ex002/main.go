// Faça um programa que receba 4 alturas usando um laço de repetição e realize a soma dessas alturas.

package main

import (
	"fmt"
)

func main() {
	var alturaTotal, soma float64

	for i := 1; i <= 4; i++ {

		fmt.Printf("Digite a altura %d:", i)
		fmt.Scan(&alturaTotal)

		if alturaTotal < 0 || alturaTotal == 0.0 {
			fmt.Println("Altura inválida")
			i--
			continue
		}

		soma += alturaTotal
	}
	fmt.Printf("Valor final: %.2f", soma)
}
