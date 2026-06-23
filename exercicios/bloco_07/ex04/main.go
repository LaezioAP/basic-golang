// Escreva um programa com a função “quadrado”,
//  que recebe um número inteiro e eleva ele ao quadrado,
//  retornando o resultado.

package main

import (
	"fmt"
	"math"
)

func quadrado(n1 float64) float64 {
	return math.Pow(n1, 2)
}

func main() {
	var number int

	fmt.Printf("Digite um número inteiro: ")
	fmt.Scanln(&number)

	q := quadrado(float64(number))

	fmt.Printf("%d * %d = %.0f\n", number, number, q)
}
