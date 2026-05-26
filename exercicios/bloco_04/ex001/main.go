// Faça um programa que conte quantas vezes a letra “a” aparece em uma palavra
package main

import (
	"fmt"
	"strings"
)

func main() {
	var contemPalavra string
	var achou int

	contemPalavra = "LAãáàezio Picanço"
	palavraEmLower := strings.ToLower(contemPalavra)

	for _, v := range palavraEmLower {
		if string(v) == "a" || string(v) == "ã" || string(v) == "â" || string(v) == "á" || string(v) == "à" {
			achou++
		}
	}

	total := strings.Count(palavraEmLower, "a") + strings.Count(palavraEmLower, "ã") + strings.Count(palavraEmLower, "á") + strings.Count(palavraEmLower, "à") + strings.Count(palavraEmLower, "â")

	fmt.Println("Contém:", total)
	// fmt.Println("Contém:", achou)
}
