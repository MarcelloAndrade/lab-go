package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID   int
	userNome string
	itens    []item
}

// (p pedido) - receiver
func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

// (p *pedido) - receiver
func (p *pedido) setUserNome(userNome string) {
	p.userNome = userNome
}

func main() {
	pedido := pedido{
		userID:   1,
		userNome: "João",
		itens: []item{
			item{1, 2, 1.5},
			item{2, 4, 2.5},
			item{produtoID: 3, qtde: 1, preco: 5.5},
		},
	}
	fmt.Println(pedido)
	fmt.Println(pedido.valorTotal())

	pedido.setUserNome("João Pedro")
	fmt.Println(pedido)
}
