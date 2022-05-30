package exercDia30_2

import (
	"fmt"
	"math/rand"
	"time"
)

type Usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Senha     int
}

func (user *Usuario) AdicionaInformacoes(nome, sobrenome, email string, idade int, senha int) {
	fmt.Printf("%p\n", user)
	user.Nome = nome
	user.Sobrenome = sobrenome
	user.Email = email
	user.Idade = idade
	user.Senha = senha
}

func (user *Usuario) ExibirInformacoes() {
	fmt.Printf("%s %s, com email %s, de %d anos, possui a senha %d\n", user.Nome, user.Sobrenome, user.Email, user.Idade, user.Senha)
}

func Ex1() {
	user := &Usuario{}
	fmt.Println(&user)
	user.AdicionaInformacoes("Orochi", "da Silva", "orochidasilva@email.com", 22, 1234)

	user2 := &Usuario{}
	fmt.Println(&user2)
	user2.AdicionaInformacoes("Matue", "da Silva", "matuedasilva@email.com", 23, 4321)

	user.ExibirInformacoes()
	user2.ExibirInformacoes()
}

type Produto struct {
	Nome       string
	Preco      float32
	Quantidade int
}

type Usuario2 struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []Produto
}

func (produto *Produto) NovoProduto(nome string, preco float32) {
	produto.Nome = nome
	produto.Preco = preco
	produto.Quantidade += produto.Quantidade
}

func AdicionarProduto(usuario *Usuario2, p *Produto, qtds int) {
	for i := 0; i < qtds; i++ {
		usuario.Produtos = append(usuario.Produtos, *p)
	}
}

func (usuario *Usuario2) DeletarProduto() {
}

func (usuario *Usuario2) ExibirUsuario() {
	fmt.Println(usuario.Nome)
	fmt.Println(usuario.Sobrenome)
	for _, p := range usuario.Produtos {
		fmt.Printf("Nome produto>: %s\n", p.Nome)
		fmt.Printf("Preco produto>: %f\n", p.Preco)
		fmt.Printf("Quantidade produto>: %d\n", p.Quantidade)
	}
	fmt.Println("-----------------------------------")

}

func Ex2() {
	user := &Usuario2{
		Nome:      "Joao",
		Sobrenome: "Belmonte",
		Email:     "joaobelmonte@email.com",
		Produtos:  []Produto{},
	}
	user2 := &Usuario2{
		Nome:      "Orochi",
		Sobrenome: "da Silva",
		Email:     "orochidasilva@email.com",
		Produtos:  []Produto{},
	}

	prod := &Produto{}
	prod.NovoProduto("Pipoca", 10.0)

	prod2 := &Produto{}
	prod2.NovoProduto("Arroz", 20.0)

	AdicionarProduto(user, prod, 3)
	AdicionarProduto(user, prod2, 5)
	AdicionarProduto(user2, prod, 2)

	user.ExibirUsuario()
	user2.ExibirUsuario()
}

type Produtos struct {
	nome       string
	preco      float32
	quantidade int
}

type Servicos struct {
	nome    string
	preco   float32
	minTrab float32
}

type Manutencao struct {
	nome  string
	preco float32
}

func somaProdutos(prods []Produtos, resultadoProdutos chan float32) {
	var somaTotal float32 = 0.0
	for _, p := range prods {
		fmt.Printf("Nome produto: %s, Preço: %.2f, Quantidade: %d\n", p.nome, p.preco, p.quantidade)
		somaTotal += (p.preco * float32(p.quantidade))
	}
	fmt.Printf("Soma total produtos: %.2f \n", somaTotal)

	resultadoProdutos <- somaTotal
}

func somaServicos(servs []Servicos, resultadoServicos chan float32) {
	var somaTotal float32 = 0.0
	for _, s := range servs {
		fmt.Printf("Nome serviço: %s, Preço: %.2f, Minutos Trabalhados : %f\n", s.nome, s.preco, s.minTrab)
		if s.minTrab < 30 {
			s.minTrab = 30
		}
		somaTotal += (s.minTrab / 60) * s.preco
	}
	fmt.Printf("Soma total serviços: %.2f \n", somaTotal)

	resultadoServicos <- somaTotal
}

func somaManutenao(manutencoes []Manutencao, resultadoManutencao chan float32) {
	var somaTotal float32 = 0.0
	for _, m := range manutencoes {
		fmt.Printf("Nome manutenção: %s, Preço: %.2f\n", m.nome, m.preco)
		somaTotal += m.preco
	}
	fmt.Printf("Soma total manutenção: %.2f \n", somaTotal)

	resultadoManutencao <- somaTotal
}

func Ex3() {
	prod1 := &Produtos{"Celular", 2000.0, 10}
	prod2 := &Produtos{"Computador", 5000.0, 4}
	prod3 := &Produtos{"Monitor", 1000.0, 12}

	serv1 := &Servicos{"Serviços 1", 1000.0, 40.0}
	serv2 := &Servicos{"Serviços 2", 3000.0, 20.0}
	serv3 := &Servicos{"Serviços 3", 2000.0, 25.0}

	man1 := &Manutencao{"Manutenção 1", 2000.0}
	man2 := &Manutencao{"Manutenção 2", 6000.0}

	arrayProdutos := []Produtos{*prod1, *prod2, *prod3}
	arrayServicos := []Servicos{*serv1, *serv2, *serv3}
	arrayManutencao := []Manutencao{*man1, *man2}

	resultadoProdutos := make(chan float32)
	resultadoServicos := make(chan float32)
	resultadoManutencao := make(chan float32)

	go somaProdutos(arrayProdutos, resultadoProdutos)
	go somaServicos(arrayServicos, resultadoServicos)
	go somaManutenao(arrayManutencao, resultadoManutencao)

	somaProd := <-resultadoProdutos
	somaServ := <-resultadoServicos
	somaMan := <-resultadoManutencao

	somaTotal := somaProd + somaServ + somaMan

	fmt.Printf("A soma dos produtos é de %.2f\n", somaProd)
	fmt.Printf("A soma dos serviços é de %.2f\n", somaServ)
	fmt.Printf("A soma das manutenções é de %.2f\n", somaMan)

	fmt.Printf("A SOMA TOTAL É DE %.2f\n", somaTotal)

	close(resultadoProdutos)
	close(resultadoServicos)
	close(resultadoManutencao)
}

func insertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		atual := array[i]
		j := i - 1
		for j >= 0 && array[j] > atual {
			array[j+1] = array[j]
			j -= 1
		}
		array[j+1] = atual
	}
}
func selectionSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		posAtual := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[posAtual] {
				posAtual = j
			}
		}
		aux := array[i]
		array[i] = array[posAtual]
		array[posAtual] = aux
	}
}

func imprimeArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func Ex4() {

	array1 := rand.Perm(100)
	t := time.Now()
	insertionSort(array1)
	fmt.Println("INSERTION: ")
	imprimeArray(array1)
	fmt.Println("Duration until t:", time.Until(t))

	array1 = rand.Perm(100)
	selectionSort(array1)
	fmt.Println("SELECT: ")
	imprimeArray(array1)
	fmt.Println("Duration until t:", time.Until(t))

	array2 := rand.Perm(1000)
	insertionSort(array2)
	fmt.Println("GROUP")
	imprimeArray(array2)
	fmt.Println("Duration until t:", time.Until(t))

	array2 = rand.Perm(1000)
	selectionSort(array2)
	fmt.Println("SELECT: ")
	imprimeArray(array2)
	fmt.Println("Duration until t:", time.Until(t))

	array3 := rand.Perm(10000)
	insertionSort(array3)
	fmt.Println("GROUP")
	imprimeArray(array3)
	fmt.Println("Duration until t:", time.Until(t))

	array3 = rand.Perm(10000)
	selectionSort(array3)
	fmt.Println("SELECT: ")
	imprimeArray(array3)
	fmt.Println("Duration until t:", time.Until(t))

}
