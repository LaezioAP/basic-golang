package main

import (
	"fmt"
)

func main() {
	nome := "Lazio Picanço"

	for i, letter := range nome {
		fmt.Println(i, letter, string(letter))
	}
}
