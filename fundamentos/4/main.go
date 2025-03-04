package main

import (
	"fmt"
)

const a = "Hello World"

type ID int
//é posswivel criar tipos

var (
	b bool    = true
	c int     = 10
	d string  = "Érick"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", f)  //  imprime o tipo
	fmt.Printf("O valor de E é %v", e) // imprime o valor
}
