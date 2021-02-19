package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobreNome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	fmt.Println("1")
	return p.nome + " " + p.sobreNome
}

func (p produto) toString() string {
	fmt.Println("2")
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimer(x imprimivel) {
	fmt.Println("3")
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimer(coisa)

	coisa = produto{"Calça", 79.50}
	fmt.Println(coisa.toString())
	imprimer(coisa)
}
