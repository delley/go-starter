package main

import (
	"fmt"

	"github.com/delley/go-starter/fundamentos/02_pacotes/operadora"
	"github.com/delley/go-starter/fundamentos/02_pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuario do sistema
var NomeDoUsuario = "Francisco"

func main() {
	fmt.Printf("Nome do usuario   : %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora : %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste    : %s\r\n", prefixo.TesteComPrefixo)
}
