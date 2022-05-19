package exercTarde

import (
	"fmt"
)

func Ex1() {
	var palavra string

	fmt.Printf("Digite uma palavra: ")
	fmt.Scanln(&palavra)

	fmt.Println("A palavra contém", len(palavra), " letras")

	for i := 0; i < len(palavra); i++ {
		fmt.Println(palavra[i])
	}
}

func Ex2(nome string, idade int8, empregado bool, umAnoDeAtividade bool, salario float32) {
	if !(idade > 22 && empregado && umAnoDeAtividade) {
		fmt.Println(nome, "não está apto a receber empréstimo")
	} else {
		if salario > 100000 {
			fmt.Println(nome, "está apto a receber empréstimo SEM juros")
		} else {
			fmt.Println(nome, "está apto a receber empréstimo COM juros")
		}
	}
}

func Ex3(mes int) {
	var meses = map[int]string{}
	meses[1] = "Janeiro"
	meses[2] = "Fevereiro"
	meses[3] = "Março"
	meses[4] = "Abril"
	meses[5] = "Maio"
	meses[6] = "Junho"
	meses[7] = "Julho"
	meses[8] = "Agosto"
	meses[9] = "Setembro"
	meses[10] = "Outubro"
	meses[11] = "Novembro"
	meses[12] = "Dezembro"

	fmt.Printf("O mês %d é o mês: %s", mes, meses[mes])
}

func Ex4() {
	var employees = map[string]int{"Benjamin": 20,
		"Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("Idade de Benjamin é", employees["Benjamin"])
	delete(employees, "Pedro")
	employees["Federico"] = 25

	var i int8 = 0
	for _, element := range employees {
		if element > 21 {
			i++
		}
	}
	fmt.Println(i, "funcionários têm mais de 21 anos.")

	fmt.Println(employees)
}
