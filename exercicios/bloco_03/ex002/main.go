package main

import (
	"fmt"
)

func main() {
	var opcao, qtdAgua int

	var valorAguaSemGas = 1.50
	var valorAguaComGas = 2.50
	fmt.Println("Bem vindo ao BeberAgua")
	fmt.Print("1 - Agua Mineral sem gás R$1,50\n")
	fmt.Print("2 - Agua Mineral com gás R$2,50\n")
	fmt.Print("Escolha uma opção: ")

	_, err := fmt.Scan(&opcao)
	if err != nil {
		fmt.Println("Valor Inválido, favor, tente novamente!")
		return
	}

	if opcao < 1 || opcao > 2 {
		fmt.Println("Produto não encontrado!")
		return
	}

	_, err = fmt.Scan(&qtdAgua)
	if err != nil {
		fmt.Println("Valor Inválido, favor, tente novamente!")
		return
	}

	if qtdAgua <= 0 {
		fmt.Println("Quantidade inválida!")
		return
	}

	switch opcao {
	case 1:
		fmt.Printf("Você escolheu Agua Mineral Sem Gás\nVALOR FINAL: R$ %.2f\n", valorAguaSemGas*float64(qtdAgua))
	case 2:
		fmt.Printf("Você escolheu Agua Mineral Com Gás\nVALOR FINAL: R$ %.2f\n", valorAguaComGas*float64(qtdAgua))
	}
}
