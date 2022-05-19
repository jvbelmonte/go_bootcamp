package exercManha

import (
	"fmt"
)

func Ex1() {
	var (
		nome  string = "João Vitor Belmonte"
		idade int    = 22
	)
	fmt.Println(nome)
	fmt.Println(idade)
}

func Ex2() {
	var (
		temperatura int8  = 17
		umidade     int8  = 72
		pressao     int16 = 1019
	)
	fmt.Println(temperatura, "°C")
	fmt.Println(umidade, "%")
	fmt.Println(pressao, "hPa")
}

func Ex3() {
	var nome string
	var sobrenome string
	var idade int
	var estatura_da_pessoa int
	var licenca_para_dirigir bool = true
	sobrenome = "6"
	quantidadeDeFilhos := 2

	fmt.Println(nome, sobrenome, idade, estatura_da_pessoa, licenca_para_dirigir, quantidadeDeFilhos)
}

func Ex4() {
	var nome string = "Fellipe"
	var sobrenome string = "Silva"
	var idade int = 25
	boolean := false
	var salario float32 = 4585.90

	fmt.Println(nome, sobrenome, idade, boolean, salario)
}
