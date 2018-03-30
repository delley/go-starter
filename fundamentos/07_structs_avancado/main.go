package main

import (
	"encoding/json"
	"fmt"

	"github.com/delley/go-starter/fundamentos/07_structs_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da vovó"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)
	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O valor da casa é: %d\r\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON:", string(objJSON))
}
