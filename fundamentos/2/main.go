package main

const a = "Hello World"
// imutável

// var b bool = false
// var c int

var (
	b bool = true
	c int
	d string
	e float64
)
// consigo alterar o valor

func main() {
	// var a string
	// ou
	a := "x"
	// : indica que é a primeira atribuição

	println(a, c, b, d, e)
}

func x() {

}

// - Já coloca valores iniciais às variáveis
// - Escopo global x escopo local
// - const imutável e var mutável
// - Variáveis em escopo local que não são usadas  causam erro, mesma coisa com imports
