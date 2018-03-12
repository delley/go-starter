package prefixo

import (
	"strconv"
)

//Capital representa o numero do prefixo de telefone da capital de um estado/província
var Capital = 11

var teste = "teste"

//TesteComPrefixo isto é só um teste
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
