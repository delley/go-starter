package operadora

import (
	"strconv"

	"github.com/delley/go-starter/fundamentos/02_pacotes/prefixo"
)

//NomeOperadora nome da operadora
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
