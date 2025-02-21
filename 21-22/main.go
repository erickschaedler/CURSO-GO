package main

import (
	"fmt"
	"github.com/erickschaedler/CURSO-GO/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	
	fmt.Println("Resultado: ", soma)
}
