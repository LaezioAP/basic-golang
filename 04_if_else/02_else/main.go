package main

import "fmt"

func main() {
	var idade int
	println("Entre com sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		println("beba a vontade!")
	} else {
		println("Não pode beber")
	}
}
