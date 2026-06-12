package main

import "fmt"

func main() {
	idade := make(map[string]int)

	idade["Laezio"] = 28

	fmt.Println("Idade:", idade)

	altura := map[string]float64{}
	altura["Laezio"] = 1.69
	fmt.Println("Altura:", altura)

	if value, ok := altura["joice"]; ok {
		fmt.Println("A altura da Joice é:", value)
	} else {
		fmt.Println("Não encontrei a altura da Joice!")
	}

}
