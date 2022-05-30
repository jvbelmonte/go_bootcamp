package exercTarde2

import (
	"fmt"
	"time"
)

type aluno struct {
	Nome         string
	Sobrenome    string
	RG           string
	DataAdmissao time.Time
}

func (a aluno) detalhes() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Sobrenome:", a.Sobrenome)
	fmt.Println("RG:", a.RG)
	fmt.Println("Data de admissão:", a.DataAdmissao)
}

func Ex1() {
	a1 := aluno{
		Nome:         "João",
		Sobrenome:    "da Silva",
		RG:           "123456",
		DataAdmissao: time.Date(1980, 01, 15, 0, 0, 0, 0, time.UTC),
	}
	a1.detalhes()
}

// Exercício 2 - Produtos de e-commerce
// Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar
// produtos e devolver o valor do preço total.
// As empresas têm 3 tipos de produtos:
// - Pequeno, Médio e Grande.
// Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

// Lista de custos adicionais:
// - Pequeno: O custo do produto (sem custo adicional)
// - Médio: O custo do produto + 3% pela disponibilidade no estoque
// - Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
// adicional pelo envio de $2500.

// 2

// Requisitos:
// - Criar uma estrutura “loja” que guarde uma lista de produtos.
// - Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
// - Criar uma interface “Produto” que possua o método “CalcularCusto”
// - Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.

// - Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome
// e preço, e devolva um Produto.

// - Será necessário uma função “novaLoja” que retorne um Ecommerce.

// - Interface Produto:
// - Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
// custo adicional segundo o tipo de produto.

// - Interface Ecommerce:
// - Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
// base no custo total dos produtos + o adicional citado anteriormente (caso a categoria
// tenha)
// - Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
// e adicioná-lo a lista da loja
type loja struct {
	listaProdutos []produto
}

type produto struct {
	tipoProduto string
	nome        string
	preco       float32
}

type Produto interface {
	CalcularCusto()
}

type Ecommerce interface {
	Total()
	Adicionar(produto produto)
}

func novoProduto(tipoProduto string, nome string, preco float32) produto {
	return produto{
		tipoProduto: tipoProduto,
		nome:        nome,
		preco:       preco,
	}
}

func novaLoja() loja {
	l := loja{
		listaProdutos: nil,
	}
	return l
}

func (p produto) CalcularCusto() float32 {
	if p.tipoProduto == "pequeno" {
		return p.preco
	}
	if p.tipoProduto == "medio" {
		return p.preco + (p.preco * 0.03)
	}
	if p.tipoProduto == "grande" {
		return p.preco + (p.preco * 0.06) + 2500
	}
	return 0.0
}

func (l loja) Total() float32 {
	var valorTotal float32 = 0.0
	for _, produto := range l.listaProdutos {
		valorTotal += produto.CalcularCusto()
	}
	return valorTotal
}

func (l loja) Adicionar(p produto) {
	l.listaProdutos = append(l.listaProdutos, p)
}

func Ex2() {

}
