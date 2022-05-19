package main

import (
	"fmt"
)

func main() {
	var i int

	fmt.Println("Entrar com o numero do exercicio: ")
	fmt.Scanln(&i)

	switch i {
	case 1:
		ex1()
	case 2:
		ex2()
	case 3:
		ex3()
	case 4:
		ex4()
	default:
		fmt.Println("Não existe exercicio com esse número")
	}
}

func ex1() {
	var (
		nome  string = "João Vitor Belmonte"
		idade int    = 22
	)
	fmt.Println(nome)
	fmt.Println(idade)
}

func ex2() {
	var (
		temperatura int8  = 17
		umidade     int8  = 72
		pressao     int16 = 1019
	)
	fmt.Println(temperatura, "°C")
	fmt.Println(umidade, "%")
	fmt.Println(pressao, "hPa")
}

func ex3() {
	var nome string
	var sobrenome string
	var idade int
	var estatura_da_pessoa int
	var licenca_para_dirigir bool = true
	sobrenome = "6"
	quantidadeDeFilhos := 2

	fmt.Println(nome, sobrenome, idade, estatura_da_pessoa, licenca_para_dirigir, quantidadeDeFilhos)
}

func ex4() {
	var nome string = "Fellipe"
	var sobrenome string = "Silva"
	var idade int = 25
	boolean := false
	var salario float32 = 4585.90

	fmt.Println(nome, sobrenome, idade, boolean, salario)
}
