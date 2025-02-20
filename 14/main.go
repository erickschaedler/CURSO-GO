package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cep        string
}

type Pessoa interface {
	DesativarCliente()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) DesativarCliente() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func DesativarPessoa(p Pessoa) {
	p.DesativarCliente()
}

func main() {
	erick := &Cliente{
		Nome:  "Wesley",
		Idade: 19,
		Ativo: true,
	}

	DesativarPessoa(erick)

	// se uma interface conter um método, todas as structs que usem aquele método
	// estão implementando a interface

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", erick.Nome, erick.Idade, erick.Ativo)
}
