package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(50, 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

// func sum(a, b int) (int, bool) {
// 	result := a+b;

// 	if result >= 50 {
// 		return result, true
// 	}

// 	return result, false
// }

func sum(a, b int) (int, error) {
	result := a + b

	if result >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return result, nil
}

//funcoes no go podem retornar mais de um valor
// nil = null
// usamos o ultimo elemento do retorno de uma funcao em go para validar os erros
