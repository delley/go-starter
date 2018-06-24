package main

import (
	"fmt"

	"github.com/delley/go-starter/fundamentos/03_funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplica, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
}
