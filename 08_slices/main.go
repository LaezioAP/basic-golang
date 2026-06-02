package main

import "fmt"

func main() {

	var x [100]int

	slice := x[:10]

	fmt.Printf("x = %T | slice = %T\n", x, slice)
	fmt.Println("Tamanho slice:", len(slice))

	y := []int{}

	fmt.Printf("y = %T\n", y)
	fmt.Println("Tamanho y:", len(y))

	for i := 1; i < 200; i++ {

	}
}
