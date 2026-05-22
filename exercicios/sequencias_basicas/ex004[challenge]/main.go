// [DESAFIO] Escreva um programa para calcular a redução do tempo de vida de um
// fumante.
// Pergunte a quantidade de cigarros fumados por dias e quantos anos ele
// já fumou. Considere que um fumante perde 10 min de vida a cada cigarro. Calcule
// quantos dias de vida um fumante perderá e exiba o total em dias.

package main

import "fmt"

func main() {
	var qtdCigarros, qtdAnosFumados int
	minPerdidoPorCigarro := 10

	fmt.Print("Quantidade de cigarro fumados por dia: ")
	fmt.Scan(&qtdCigarros)

	fmt.Print("Quantidade de anos que você já fumou: ")
	fmt.Scan(&qtdAnosFumados)

	minutosPerdidoNoDia := qtdCigarros * minPerdidoPorCigarro

	minutosPedidosNoAno := minutosPerdidoNoDia * 360

	totalDeMinutosPerdido := minutosPedidosNoAno * qtdAnosFumados

	totalDiasPerdido := totalDeMinutosPerdido / 1440

	fmt.Printf("\nVocê perderá aproximadamente %d dias de vida.\n", totalDiasPerdido)
}

// ANO 360
// 1 dia = 24h =  1440min
// 60min = 1h
