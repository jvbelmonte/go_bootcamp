package exercManha2

import (
	"errors"
	"fmt"
)

const (
	minimum   = "minimum"
	average   = "average"
	maximum   = "maximum"
	cao       = "cao"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func calculaImposto(salario float32) (float32, float32) {
	impostoPorcentagem := (salario * 27) / 150000
	impostoValor := (salario * impostoPorcentagem) / 100
	return impostoPorcentagem, impostoValor
}
func calculaMedia(quantidade, somaDasNotas int16) int16 {
	return somaDasNotas / quantidade
}
func calculaSalario(salario, horas, aditivoHoras, porcentagem int) int {

	horasExtras := aditivoHoras * porcentagem
	salario *= horas
	salario += horasExtras

	return salario
}
func getFunc(opt string) (func(valores ...float32) (float32, error), error) {
	if opt == minimum {
		return minFunc, nil
	}
	if opt == maximum {
		return maxFunc, nil
	}
	if opt == average {
		return avgFunc, nil
	}
	return nil, errors.New("Função Inválida")
}
func getFuncAnimais(nomeAnimal string) (func(quantidadeAnimal int) int, error) {
	if nomeAnimal == cao {
		return cachorroFunc, nil
	}
	if nomeAnimal == gato {
		return gatoFunc, nil
	}
	if nomeAnimal == hamster {
		return hamsterFunc, nil
	}
	if nomeAnimal == tarantula {
		return tarantulaFunc, nil
	}
	return nil, errors.New("Animal Inválido")
}
func minFunc(valores ...float32) (float32, error) {
	if len(valores) == 0 {
		return 0.0, errors.New("Valor vazio")
	}
	minValue := valores[0]
	for _, valor := range valores {
		if valor < minValue {
			minValue = valor
		}
	}
	return minValue, nil
}
func maxFunc(valores ...float32) (float32, error) {
	if len(valores) == 0 {
		return 0.0, errors.New("Valor vazio")
	}
	maxValue := valores[0]
	for _, valor := range valores {
		if valor > maxValue {
			maxValue = valor
		}
	}
	return maxValue, nil
}
func avgFunc(valores ...float32) (float32, error) {
	if len(valores) == 0 {
		return 0.0, errors.New("Valor vazio")
	}
	var avgValue float32 = 0
	for _, valor := range valores {
		avgValue += valor
	}
	avgValue /= float32(len(valores))

	return avgValue, nil
}
func cachorroFunc(quantidadeCachorro int) int {
	return quantidadeCachorro * 10000
}
func gatoFunc(quantidadeGato int) int {
	return quantidadeGato * 5000
}
func hamsterFunc(quantidadeHamster int) int {
	return quantidadeHamster * 250
}
func tarantulaFunc(quantidadeTarantulas int) int {
	return quantidadeTarantulas * 150
}
func Ex1(salario float32) {
	impostoPorcentagem, impostoValor := calculaImposto(salario)
	fmt.Println("O imposto do salário com valor", salario, "é de", impostoPorcentagem, "%, o equivalente a", impostoValor)
}
func Ex2(notas ...int16) (int16, error) {
	var i, sumNotas int16 = 0, 0
	for _, nota := range notas {
		if nota == 0 {
			return 0, errors.New("Valor 0 não é aceitável")
		}
		sumNotas += nota
		i++
	}
	media := calculaMedia(i, sumNotas)

	return media, nil
}
func Ex3(categoria string, horas int) {
	restoHoras := 0
	switch categoria {
	case "A":
		if horas > 160 {
			restoHoras = horas - 160
		}
		salarioFinal := calculaSalario(3000, horas, restoHoras, 50)
		fmt.Println("Salário Final:", salarioFinal)
	case "B":
		if horas > 160 {
			restoHoras = horas - 160
		}
		salarioFinal := calculaSalario(1500, horas, restoHoras, 20)
		fmt.Println("Salário Final:", salarioFinal)
	case "C":
		salarioFinal := calculaSalario(1000, horas, restoHoras, 0)
		fmt.Println("Salário Final:", salarioFinal)
	}
}
func Ex4() {
	minFunc, _ := getFunc(minimum)
	avgFunc, _ := getFunc(average)
	maxFunc, _ := getFunc(maximum)

	min, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("max: %.2f\n", max)
}
func Ex5() {
	nomeAnimal := "hamster"

	qntdComidaG, _ := getFuncAnimais(nomeAnimal)
	fmt.Printf("quandidade de alimento do(s) %s (em gramas):	%d gramas\n", nomeAnimal, qntdComidaG(5))

}
