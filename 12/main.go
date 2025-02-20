package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cep string
}

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	Address Endereco
	// ou
	Endereco
}

func main() {
	wesley := Cliente {
		Nome: "Wesley",
		Idade: 30,
		Ativo: true,
	}
	
	wesley.Ativo = false
	wesley.Logradouro = "São Paulo"
	// herdando x criando uma chave
	wesley.Address.Logradouro = "São Paulo"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)
}