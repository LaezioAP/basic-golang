// Faça um programa que use uma função para responder
//  a saudação de um usuário: “Olá, fulano! Boas vindas!”.
// Use funções para isso.

package main

import "fmt"

func welcomeUser(name string) {
	fmt.Printf("Olá %s, seja bem vindo!", name)
}

func main() {
	var name string
	fmt.Printf("Entre com seu nome: ")
	fmt.Scanf("%s", &name)
	welcomeUser(name)
}
