// Escreva uma função que recebe a e b e troque seus valores.
// a = 1; b=2 -> a=2; b=1

package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func input() {
	var n1, n2 int

	fmt.Printf("Digite o número 1: ")
	fmt.Scanln(&n1)

	fmt.Printf("Digite o número 2: ")
	fmt.Scanln(&n2)

	fmt.Printf("Iniciou com: 1 = %d | 2 = %d\n", n1, n2)

	swap(&n1, &n2)

	fmt.Printf("Finalizou com: 1 = %d | 2 = %d", n1, n2)

}

func main() {
	input()

}
