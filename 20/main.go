package main

type MyNumber int

// constraint
type Number interface {
	// int | float64 | MyNumber
	// ou
	~int | ~float64
}

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

// para caso a funcao sirva para mais de um tipo
func Soma[T Number] (m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func main() {
	m := map[string]int{"Erick": 10000, "João": 2000}

	println(Soma(m))

	m2 := map[string]float64{"Erick": 10000.50, "João": 2000.50}

	println(Soma(m2))

	m3 := map[string]MyNumber{"Erick": 10000, "João": 200}

	println(Soma(m3))

	println(Compara(10, 10.0))
}
