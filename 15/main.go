package main

func main() {
	a := "erick"

	// ponteiro := &a

	// println(*ponteiro) erick
	// println(ponteiro)  endereco na memoria: 0x14000054728

	// nova variavel com o valor do ponteiro de a
	b := &a
	// muda o valor que é referenciado pelo ponteiro de "a" e que é "b"
	*b = "erick2"

	println(a) // erick 2
	println(b)
}

// ponteiro é o endereco do variavel na memoria