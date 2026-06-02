// Faça um programa que receba 6 temperaturas. Remova a 1a e a última para calcular a média. Se a média for acima de 30 graus, exiba que houve um aumento da temperatura.
package main

import "fmt"

func main() {
	temps := []float64{}
	var temp, mediaTemps, soma float64
	totalTemps := 6

	for i := 1; i <= totalTemps; i++ {
		fmt.Print("Registre temperaturas: ")
		fmt.Scan(&temp)

		temps = append(temps, temp)
	}

	removeFirstAndLastTemp := temps[1 : len(temps)-1]

	for _, value := range removeFirstAndLastTemp {
		soma += value
	}

	mediaTemps = soma / float64(len(removeFirstAndLastTemp))

	if mediaTemps > 30 {
		fmt.Println("Temperatura aumentada")
	} else {
		fmt.Println("Temperatura normal")
	}
}
