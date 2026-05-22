// Faça um programa que receba um número inteiro e
// calcule sua raiz quadrada e exiba o resultado.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n1 int

	fmt.Print("Digite um número: ")
	fmt.Scan(&n1)

	fmt.Printf("Raiz Quadrada: %.2f", math.Sqrt(float64(n1)))
}
