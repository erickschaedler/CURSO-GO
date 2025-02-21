package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

// para exportador do arquivo
// por regrad do Go
// qualquer coisa deve se ter letra maiuscula