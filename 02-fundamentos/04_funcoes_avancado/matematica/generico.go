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

/*
Dividi calcula a divisão de dois numeros
*/
func Dividi(x int, y int) (resultado int) {
	resultado = x / y
	return
}

/*
DividiComResto calcula a divisão de dois numeros e o resto da divisao
*/
func DividiComResto(x int, y int) (resultado int, resto int) {
	resultado = x / y
	resto = x % y
	return
}
