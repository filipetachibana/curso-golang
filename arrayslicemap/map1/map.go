package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678987] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95412385554] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95412385554])
	delete(aprovados, 95412385554)
	fmt.Printf(aprovados[95412385554])
}
