package main

import "fmt"

func init() {
	fmt.Println("Inicializando....")
}

func main() {
	i := 1

	var p *int = nil // aloca 64bytes de memória para um ponteiro vazio
	p = &i           // &i get endereço de memória
	*p++             // *p get o valor dentro do ponteiro
	i++
	fmt.Println(*p, p, &p)
	fmt.Println(i, &i)
}
