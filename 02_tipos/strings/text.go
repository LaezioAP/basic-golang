package main

import "fmt"

func main() {
	fmt.Println("Daily")
	fmt.Println(`Laezio tem`, len("Laezio"), "letras")

	fmt.Println("Laezio"[0])         // Isso representa um BYTE (int8)
	fmt.Println(string("Laezio"[0])) // Isso representa strings
}
