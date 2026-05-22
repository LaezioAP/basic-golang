// A locadora de carros precisa da sua ajuda para cobrar seus serviços. Escreva
// um programa que pergunte a quantidade de Km percorridos por um carro alugado e a
// quantidade de dias pelos quais ele foi alugado. Calcule o preço total a pagar,
// sabendo que o carro custa R$90 por dia e R$0,20 por Km rodado.

package main

import "fmt"

func main() {
	var km float64
	var dias int
	fmt.Println("Welcome to APP")

	fmt.Print("Km percorrido: ")
	fmt.Scan(&km)

	fmt.Print("Dias alugados: ")
	fmt.Scan(&dias)

	valorPerDia := float64(dias) * 90
	valorPerKm := km * 0.20

	totalPagar := valorPerDia + valorPerKm

	fmt.Printf("Valor por KM: R$%.2f\n", valorPerKm)
	fmt.Printf("Valor por DIA: R$%.2f\n", valorPerDia)
	fmt.Printf("Total a Pagar: R$%.2f", totalPagar)
}
