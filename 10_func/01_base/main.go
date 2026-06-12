package main

import "fmt"

// Função que soma valores inteiros
func soma(a int, b int) int {
	return a + b
}

func media(a, b int) (res float64, erro error) {
	total := soma(a, b)
	res = float64(total) / 2
	erro = nil
	return res, erro
}

func main() {
	n1, n2 := 10, 20
	fmt.Println(n1, n2)

	total := soma(n1, n2) // invocou a soma.
	fmt.Println("Soma dos valores A + B =", total)

	resMedia, erro := media(n1, n2)
	if erro != nil {
		fmt.Println("Deu erro:", erro)
	}
	fmt.Println("Média dos valores:", resMedia)
}
