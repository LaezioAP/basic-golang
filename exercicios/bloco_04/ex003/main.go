package main

import (
	"fmt"
	"strconv"
)

func main() {
	var soma float64

	for {
		var entrada string
		fmt.Print("Entre com o valor(digite 0 para cancelar): ")
		fmt.Scan(&entrada)

		valor, err := strconv.ParseFloat(entrada, 64)
		if err != nil {
			fmt.Print("Valor inválido")
			continue
		}

		if valor == 0 {
			break
		}

		soma += valor
	}

	fmt.Printf("Saldo total: R$%.2f", soma)
}
