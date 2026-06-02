package main

import "fmt"

// type WeekDay int

// const (
// 	Sunday WeekDay = iota
// )

func main() {
	var notas [4]float64
	var soma float64

	for i := range notas {
		fmt.Printf("\nNota %d:", i+1)
		fmt.Scan(&notas[i])
	}

	for _, v := range notas {
		soma += v
	}
	// fmt.Println(Sunday)

	fmt.Println("Nota final:", soma)
	fmt.Println("Media final:", soma/float64(len(notas)))
}
