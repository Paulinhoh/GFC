/*
Todo programa em Go tem que ter um package
que normalmente é o nome do diretorio tirando o arquivo
main que o arquivo principal do programa
*/
package main // package main

import "fmt" // importação de pacotes

// no go é possivel a criação da própria tipagem
type ID int
var f ID = 8

// variaveis globais
const b bool = true// declaração de constante

// declaração de variaveis em bloco
var (
	c int = 10
	d string
	e float64
)


func main() {
	println("Hello, World!")
	
	a := 10 // variavel local e declaração de variavel curta
	println(a)

	fmt.Printf("O tipo de E é: %T\n", e)

	// =-=-=-=-=-= Trabalhando com Arrays -=-=-=-=-=-=-=
	var meuArray [3]int // declaração de array
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray))
	fmt.Println(meuArray[2])

	for indice, valor := range meuArray {
		fmt.Printf("Indice: %d, Valor: %d\n", indice, valor)
	}

	// =-=-=-=-=-= Trabalhando com Slices -=-=-=-=-=-=-=
	meuSlice := []int{10, 20, 30, 40, 50} // pode ser inicializado vazio
	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)
	
	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice[:0]), cap(meuSlice[:0]), meuSlice[:0])
	
	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice[:4]), cap(meuSlice[:4]), meuSlice[:4])
	
	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice[2:]), cap(meuSlice[2:]), meuSlice[2:])
	
	meuSlice = append(meuSlice, 60)
	fmt.Printf("len=%d cap=%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)
}
