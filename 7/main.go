package main

import "fmt"

func main() {
	salarios := map[string]int{"Erick": 100000, "Joao": 2000, "Maria": 3000}
	// cria com valores

	// fmt.Println(salarios["Erick"])

	delete(salarios, "Joao")
	
	salarios["Pedro"] = 5000

	// fmt.Printf("%v\n", salarios)

	// sal := make(map[string]int)
	// sa1l1:= map[string]int{}
	// cria sem valores

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}

// array = tem tamanho fixo definido na criacao
// slice = array de elementos sem tamanho fixo
// map = array de chave valor
