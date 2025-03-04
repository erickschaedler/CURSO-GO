package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20
	println(soma(&var1, &var2)) // 50
	println(var1) // 50
}

// passar o endereco da variavel na memoria por parametro para atualizar o valor
// dentro de uma funcao secundaria e refletir na principal

//  para trabalhar com valores mutaveis
