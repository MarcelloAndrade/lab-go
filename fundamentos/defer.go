package main

import "fmt"

func main() {
	defer fmt.Println("Fim")
	fmt.Println("Inicio")
}
