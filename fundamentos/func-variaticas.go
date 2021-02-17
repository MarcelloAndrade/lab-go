package main

import "fmt"

func main() {
	approved := []string{"Pedro", "Maria", "João"} // slice
	printApproved(approved...)                     // ... separa os valores como se estivesse passando "Pedro", "Maria", "João"
}

func printApproved(approved ...string) {
	for i, ap := range approved {
		fmt.Printf("%d) %s\n", i+1, ap)
	}
}
