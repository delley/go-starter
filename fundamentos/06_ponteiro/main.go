package main

import "fmt"

//Imovel eh uma struct que armazena um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel)
	fmt.Printf("A casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara linda", 350000}
	fmt.Printf("A chacara é: %p - %+v\r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("A chacara é: %p - %+v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
