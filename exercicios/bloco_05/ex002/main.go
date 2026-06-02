// Faça um programa que receba 4 notas, calcule a média, mínimo e máximo dessas notas.

package main

import (
	"fmt"
	"slices"
)

func main() {
	var nota, soma, media float64
	notas := []float64{}

	for i := 0; i < 5; i++ {
		fmt.Printf("Nota %d:", i+1)
		fmt.Scan(&nota)

		notas = append(notas, nota)
	}

	for _, valor := range notas {
		soma += valor
	}

	media = soma / float64(len(notas))

	fmt.Println("Média: ", media)
	fmt.Println("Max: ", slices.Max(notas))
	fmt.Println("Min: ", slices.Min(notas))
}
