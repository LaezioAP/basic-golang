package main

import "fmt"

func main() {

	var opcao int
	fmt.Print("Digite um opção: (1,2,3,4,5)")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		fmt.Println("Maçã")
	case 2:
		fmt.Println("Banana")
	case 3:
		fmt.Println("Pêra")
	case 4:
		fmt.Println("Uva")
	default:
		fmt.Println("Não há registros...")
	}
}
