package main

import (
	"errors"
	"fmt"
)

// Escreva uma função que receba um slice de inteiros e uma função como parâmetros. A função deve aplicar a função recebida em cada elemento do slice e retornar a soma de todos os resultados. Caso o slice esteja vazio, retorne um erro.
func applyFunctionToSlice(slice []int, function func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Slice vazio")
	}

	sum := 0
	for _, value := range slice {
		result := function(value)
		sum += result
	}

	return sum, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	square := func(num int) int {
		return num * num
	}

	result, err := applyFunctionToSlice(slice, square)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado:", result)
}
