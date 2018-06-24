package main

import (
	"fmt"
)

var (
	//Endereco é um valor importante e tem que ser publico
	Endereco string
	//Telefone é um valor importante para nossa aula
	Telefone   string
	quantidade int     // quantidade = 0
	comprou    bool    // comprou = false
	valor      float64 // valor = 0.00
	palavras   rune
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("Endereço: %s\n\r", Endereco)
	fmt.Printf("Quantidade: %d\n\r", quantidade)
	fmt.Printf("Comprou: %v\n\r", comprou)
	fmt.Printf("Palavras: %v\n\r", palavras)
	fmt.Printf("O valor de teste é: %s\n\r", teste)
}
