package exercDia30_1

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type produto struct {
	idProduto  string
	preco      float32
	quantidade int
}

func (p produto) retornaLinha() string {
	return fmt.Sprintf("%s, %f, %d\n", p.idProduto, p.preco, p.quantidade)
}

func geraCSV(caminho string, produtos []produto) error {

	if len(produtos) == 0 {
		return errors.New("Lista zerada!")
	}

	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0600)

	defer file.Close()

	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}

	//p := produtos[0]

	if _, err = file.WriteString("id,preco,quantidade\n"); err != nil {
		return fmt.Errorf("erro ao gerar cabeçalho: %w", err)
	}

	for _, prod := range produtos {
		if _, err = file.WriteString(prod.retornaLinha()); err != nil {
			return fmt.Errorf("error ao salvar linha: %w", err)
		}
	}

	return nil
}

func lerCSV(caminho string) error {

	file, err := os.Open(caminho)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := tabwriter.NewWriter(os.Stdout, 20, 30, 1, '\t', tabwriter.AlignRight)

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	cabecalho := strings.Split(scanner.Text(), ",")

	for _, c := range cabecalho {
		fmt.Fprintf(w, "%s\t", c)
	}

	fmt.Fprintln(w)

	var sum float64 = 0.0
	var count int = 1
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for index, v := range values {
			if index == 1 || index == count {
				v := strings.ReplaceAll(v, " ", "")
				conv, _ := strconv.ParseFloat(v, 64)
				sum += conv
				count += count + 2
			}
			fmt.Fprintf(w, "%s\t", v)
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintf(w, "Soma Total:\n%f\t", sum)
	fmt.Fprintln(w)

	w.Flush()

	return nil
}

func Ex1() {
	produtos := []produto{
		{
			idProduto:  "3",
			quantidade: 8,
			preco:      6.99,
		},
		{
			idProduto:  "2",
			quantidade: 20,
			preco:      12.99,
		},
		{
			idProduto:  "1",
			quantidade: 10,
			preco:      9.99,
		},
	}
	geraCSV("produtos.csv", produtos)
}

func Ex2() {
	lerCSV("produtos.csv")
}
