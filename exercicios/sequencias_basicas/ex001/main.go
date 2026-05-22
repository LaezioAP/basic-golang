// Faça um programa que leia as duas notas de um aluno em uma matéria e mostre
// na tela a sua média na disciplina.
package main

import "fmt"

func main() {
	var v1, v2 float64

	fmt.Print("Nota 1: ")
	fmt.Scan(&v1)

	fmt.Print("Nota 2: ")
	fmt.Scan(&v2)

	media := (v1 + v2) / 2

	fmt.Printf("Média: %v", media)
}
