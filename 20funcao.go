package main

import (
	"errors"
	"fmt"
)

// Escreva uma função que receba um slice de strings como parâmetro e retorne um novo slice contendo apenas as strings que possuem mais de 5 caracteres. Caso o slice esteja vazio, retorne um erro.
func filterLongStrings(strings []string) ([]string, error) {
	if len(strings) == 0 {
		return nil, errors.New("slice vazio")
	}

	var longStrings []string
	for _, s := range strings {
		if len(s) > 5 {
			longStrings = append(longStrings, s)
		}
	}

	return longStrings, nil
}

func main() {
	strings := []string{"gato", "cachorro", "elefante", "pássaro", "formiga", "leão"}
	filteredStrings, err := filterLongStrings(strings)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(filteredStrings)
}
