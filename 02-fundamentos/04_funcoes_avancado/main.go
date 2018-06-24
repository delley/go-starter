package main

import (
	"fmt"

	"github.com/delley/go-starter/fundamentos/04_funcoes_avancado/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplica, 2, 2)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Dividi, 12, 3)
	fmt.Printf("O resultado da divisão foi: %d\r\n", resultado)
	resultado, resto := matematica.DividiComResto(12, 5)
	fmt.Printf("O resultado da divisão foi: %d e o resto foi: %d\r\n", resultado, resto)
}
