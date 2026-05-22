package main

import "fmt"

func main() {
	var nota, soma float64

	fmt.Print("Nota 1: ")
	fmt.Scan(&nota)
	soma += nota

	fmt.Print("Nota 2: ")
	fmt.Scan(&nota)
	soma += nota

	fmt.Print("Nota 3: ")
	fmt.Scan(&nota)
	soma += nota

	media := soma / 3

	if media >= 6 {
		fmt.Printf("Você passou! Média = %.2f\n", media)
	} else {
		fmt.Printf("Reprovadoooo! Média = %.2f\n", media)
	}
}
