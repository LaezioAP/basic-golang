// Faça um programa que calcule a média de uma quantidade indefinida
// de valores, usando uma função para isso.

package main

import (
	"fmt"
	"strconv"
)

func getValues() []float64 {
	valores := []float64{}
	var valorStr string

	for {
		fmt.Printf("Digite os valores que deseja validar a media: ")
		_, err := fmt.Scanln(&valorStr)

		if err != nil {
			if err.Error() == "unexpected newline" {
				return valores
			} else {
				fmt.Println("Favor, inserir 1 número de cada vez!")

				var lixo string
				fmt.Scanln(&lixo)

				continue
			}
		}

		valor, err := strconv.ParseFloat(valorStr, 64)
		if err != nil {
			fmt.Println("Valor inválido!")
			continue
		}

		valores = append(valores, valor)
	}
}

func somarValores(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}

	return sum
}

func calcMedia(values []float64) float64 {
	if len(values) == 0 {
		return 0.0
	}

	return somarValores(values) / float64(len(values))
}

func main() {
	values := getValues()
	mediaValor := calcMedia(values)

	fmt.Println("Valores: ", values)
	fmt.Println("Soma dos valores: ", mediaValor)
}
