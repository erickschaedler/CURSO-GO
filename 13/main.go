package main

import "fmt"

type Endereco struct {
  Logradouro string
  Numero     int
  Cep        string
}

type Cliente struct {
  Nome    string
  Idade   int
  Ativo   bool
  Endereco
}

func (c *Cliente) DesativarCliente() {
  c.Ativo = false
  fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
  erick := Cliente{
    Nome:  "Wesley",
    Idade: 19,
    Ativo: true,
  }

  erick.DesativarCliente()

  fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", erick.Nome, erick.Idade, erick.Ativo)
}