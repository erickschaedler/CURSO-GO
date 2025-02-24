package main

func main() {
	a := 1
	b := 2
	c := 3

	switch a {
	case 1:
		println(1)
	case 2:
		println(2)
	default:
		println(3)
	}

	if a > b && c > a {
		// para ou usa-se ||
		println(c)
	}

	if a > b {
		println(a)
	} else {
		println(b)
	}
}

// no go é possivel escolher o ambiente de execucao
// atravez da variavel de ambiente GOOS

// GOOS=windows go build main.go

// go env mostra as variaveis de ambiente

// GOOS=windows GOARCH=amd64 go build main.go

// go tool dist list montra todas as opcoes

// go env GOOS GOARCH mostra o padrao do computador atual

// para gerar o nome do binario personalizado pode-se usar
// go build -o nome