package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)

	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos) ... uint8 uint16 uint32 uint64

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("o rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string

	s1 := "olá meu nome é Filipe"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string é", len(s1))

	//string com mutiplas linhas

	s2 := `olá
	meu
	nome
	é
	Filipe`
	fmt.Println("O tamanho da string é", len(s2))

	// char??
	// var x char = 'b'
	char := 'a'
	fmt.Println("O tipo char é", reflect.TypeOf(char))
	fmt.Println(char)

}
