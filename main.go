package main

import (
	"fmt"

	"github.com/jvbelmonte/go_bootcamp/exercManha"
	"github.com/jvbelmonte/go_bootcamp/exercTarde"
)

func main() {
	var i int

	fmt.Println("Entrar com o numero do exercicio: ")
	fmt.Scanln(&i)

	switch i {
	case 1:
		exercManha.Ex1()
	case 2:
		exercManha.Ex2()
	case 3:
		exercManha.Ex3()
	case 4:
		exercManha.Ex4()
	case 5:
		exercTarde.Ex1()
	case 6:
		exercTarde.Ex2("Joao", 23, false, true, 2000)
	case 7:
		exercTarde.Ex3(1)
	case 8:
		exercTarde.Ex4()
	default:
		fmt.Println("Não existe exercicio com esse número")
	}
}
