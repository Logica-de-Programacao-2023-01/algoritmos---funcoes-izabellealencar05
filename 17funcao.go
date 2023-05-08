package main

import (
	"fmt"
	"sort"
)

// Crie uma função que receba um slice de strings como parâmetro e retorne uma nova string com as strings em ordem alfabética. Caso o slice esteja vazio, retorne um erro.
func ordenarStrings(s []string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("O slice está vazio")
	}

	sort.Strings(s)
	return fmt.Sprintf("%v", s), nil
}

func main() {
	s := []string{"banana", "maçã", "abacaxi", "laranja"}
	resultado, err := ordenarStrings(s)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
