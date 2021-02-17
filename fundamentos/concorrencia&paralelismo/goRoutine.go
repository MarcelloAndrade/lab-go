// Paralelismo - executa código simultaneamente em processadores físicos diferentes.

// Concorrencia - intercala (administra) vários processos ao mesmo tempo e isso pode ocorrer em um único processador físico
package main

import (
	"fmt"
	"time"
)

func main() {
	go fale("Maria", "Oi ..", 500)
	go fale("João", "Adeus ..", 500)
	fale("Pedro", "Estou aqui ..", 5)
}

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}
