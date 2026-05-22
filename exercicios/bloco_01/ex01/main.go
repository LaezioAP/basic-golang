package main

import "fmt"

func main() {
	celsiu := 40

	fahren := (celsiu * 9 / 5) + 32

	fmt.Printf("Celsius: %v. Convertido para Fahrenheit: %v", celsiu, fahren)
}
