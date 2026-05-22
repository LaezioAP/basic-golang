package main

import "fmt"

func main() {
	fmt.Print("Digite seu nome: ")

	var name string
	fmt.Scanf("%s", &name)

	fmt.Printf("Seja bem vindo, %v", name)
}
