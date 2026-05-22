package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	leitor := bufio.NewReader(os.Stdin)
	fmt.Print("Digite seu nome: ")

	name, _ := leitor.ReadString('\n')

	name = strings.TrimSpace(name)
	valorEmLower := strings.ToLower(name)

	if strings.Contains(valorEmLower, "calvo") {
		fmt.Println("Alerta: Identificamos um calvo no sistema!")
	} else {
		fmt.Println("Cabelo detectado. Nome limpo.")
	}
}
