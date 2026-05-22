package main

import "fmt"

func main() {
	var idade int
	println("Entre com sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 66 {
		println("Cuidado com a bebida!")
	} else if idade >= 18 {
		println("Beba a vontade")
	} else {
		println("Não pode beber")
	}
}
