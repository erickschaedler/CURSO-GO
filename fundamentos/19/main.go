package main

import "fmt"

func main() {
	var minhaVar interface{} = "erick"
	println(minhaVar.(string)) // define o tipo
	res, ok := minhaVar.(int)
	// sem o ok ele da um panic
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}

//type assertion
