package main

import (
	"fmt"
)

func main() {
	meses := 6
	situacao := true
	cidade := "Teste"

	// Operadores: < > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve a pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele está devendo.")
	}

	if !situacao {
		fmt.Println("Ele está adimplente.")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive no estado Teste.")
	}

	if descricao, status := checaSituacaoCliente(meses); status {
		fmt.Println("Qual a situação do cliente?", descricao)
	}
}

func checaSituacaoCliente(meses int) (descricao string, status bool) {
	descricao = "Adimplente"
	if meses > 0 {
		status = true
		descricao = "Inadimplente"
	}
	return
}
