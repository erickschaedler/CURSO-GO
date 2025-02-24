package main

import (
	"curso-go/matematica"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(10, 20)
	
	fmt.Println("Resultado: ", soma)
	fmt.Println(uuid.New())
}
