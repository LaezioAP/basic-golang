// Funções variaticas
package main

import "fmt"

func soma(values ...int) int {

	total := 0
	for _, v := range values {
		total += v
	}

	return total
}

func main() {
	var a, b int

	a = 10
	b = 20
	// c := 30

	total := soma(a, b)
	fmt.Println(total)

	valores := []int{1, 2, 3, 4, 5, 6, 10, 20, 40}
	total = soma(valores...)
	fmt.Println(total)
}
