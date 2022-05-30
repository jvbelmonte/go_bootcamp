package main

import (
	"fmt"

	exercManha1 "github.com/jvbelmonte/go_bootcamp/dia1/exercManha"
	exercTarde1 "github.com/jvbelmonte/go_bootcamp/dia1/exercTarde"
	exercManha2 "github.com/jvbelmonte/go_bootcamp/dia2/exercManha"
	exercTarde2 "github.com/jvbelmonte/go_bootcamp/dia2/exercTarde"
	exercDia30_1 "github.com/jvbelmonte/go_bootcamp/dia3/dia01"
	exercDia30_2 "github.com/jvbelmonte/go_bootcamp/dia3/dia02"
)

func main() {
	var i int

	fmt.Println("Entrar com o numero do exercicio: ")
	fmt.Scanln(&i)

	switch i {
	case 1:
		exercManha1.Ex1()
	case 2:
		exercManha1.Ex2()
	case 3:
		exercManha1.Ex3()
	case 4:
		exercManha1.Ex4()
	case 5:
		exercTarde1.Ex1()
	case 6:
		exercTarde1.Ex2("Joao", 23, false, true, 2000)
	case 7:
		exercTarde1.Ex3(1)
	case 8:
		exercTarde1.Ex4()
	case 9:
		exercManha2.Ex1(50000)
	case 10:
		media, err := exercManha2.Ex2(0, 5, 10, 9, 1)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Media dos alunos é", media)
	case 11:
		exercManha2.Ex3("A", 120)
	case 12:
		exercManha2.Ex4()
	case 13:
		exercManha2.Ex5()
	case 14:
		exercTarde2.Ex1()
	case 15:
		exercTarde2.Ex2()
	case 16:
		exercDia30_1.Ex1()
	case 17:
		exercDia30_1.Ex2()
	case 18:
		exercDia30_2.Ex1()
	case 19:
		exercDia30_2.Ex2()
	case 20:
		exercDia30_2.Ex3()
	case 21:
		exercDia30_2.Ex4()
	case 0:
		break
	default:
		fmt.Println("Não existe exercicio com esse número")
	}
}
