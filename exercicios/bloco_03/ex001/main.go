package main

import (
	"fmt"
)

func main() {
	var opcao int
	fmt.Println("Bem vindo ao BeberAgua")
	fmt.Print("1 - Agua Mineral sem gás R$1,50\n")
	fmt.Print("2 - Agua Mineral com gás R$2,50\n")
	fmt.Print("Escolha uma opção: ")

	_, err := fmt.Scan(&opcao)
	if err != nil {
		fmt.Println("Valor Inválido, favor, tente novamente!")
		return
	}

	switch opcao {
	case 1:
		fmt.Println("Você escolheu Agua Mineral Sem Gás - R$1,50")
	case 2:
		fmt.Println("Você escolheu Agua Mineral Com Gás - R$2,50")
	default:
		fmt.Println("Sem opção!")
	}
}
