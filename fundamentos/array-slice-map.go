package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ----------- ----------- ----------- //
	fmt.Println("## For")
	var notas [3]int
	notas[0], notas[1], notas[2] = 1, 2, 3
	fmt.Println(notas)

	for i := 0; i < len(notas); i++ {
		fmt.Println("Nota:", notas[i])
	}

	// ----------- ----------- ----------- //
	fmt.Println("## For range")
	numeros := [...]int{1, 2, 3, 4, 5} // ... compilador conta o tamanho do array
	for i, numero := range numeros {
		fmt.Printf("%d) %d \n", i, numero)
	}

	// ----------- ----------- ----------- //
	fmt.Println("## Slice")
	a1 := [3]int{1, 2, 3} // array - quando definimos o tamanho
	s1 := []int{1, 2, 3}  // slice - quando não definimos o tamanho
	fmt.Println(a1, s1)
	fmt.Println("Array", reflect.TypeOf(a1), "- Slice", reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// ----------- ----------- ----------- //
	fmt.Println("## Slice Make")
	sm := make([]int, 10)
	sm[9] = 12
	fmt.Println(sm, len(sm), cap(sm))

	sm1 := make([]int, 10, 20)
	fmt.Println(sm1, len(sm1), cap(sm1))

	// ----------- ----------- ----------- //
	fmt.Println("## Map")
	//var m1 map[int]string
	m1 := make(map[int]string)
	m1[1] = "Var 1"
	m1[2] = "Var 2"
	m1[3] = "Var 3"
	fmt.Println(m1)
	fmt.Println(m1[1])
	delete(m1, 1)
	fmt.Println(m1)

	for x, y := range m1 {
		fmt.Println(x, y)
	}

	fmt.Println(teste())
}

func teste() (string, string, string, string, int) {
	return "1", "2", "3", "4", 10
}
