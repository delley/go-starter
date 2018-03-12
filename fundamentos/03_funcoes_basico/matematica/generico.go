package matematica

/*
Calculo executa qualuqer tipo de calculo de acordo com a funcao passada como parâmetro
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

/*
Multiplica calcula o produto de dois numeros
*/
func Multiplica(x int, y int) int {
	return x * y
}
