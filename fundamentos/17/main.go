package main

import "fmt"

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	conta := Conta{
		saldo: 100,
	}

	conta.simular(200)

	fmt.Println(conta.saldo)
}