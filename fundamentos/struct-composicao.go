package main

import "fmt"

type carro struct {
	nome       string
	velocidade float64
}

type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidade = 0
	f.turboLigado = false

	fmt.Println(f)
}
