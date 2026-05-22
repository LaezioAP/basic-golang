package main

import "fmt"

func main() {
	var idade int
	println("Entre com sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		println("beba a vontade!")
	}

	if idade <= 17 {
		println("Não pode beber")
	}
}
