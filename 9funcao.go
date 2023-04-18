package main

import (
	"fmt"
	"strings"
)

// Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string. Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro.
func split(s string) ([]string, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("a sring esta vazia")
	}
	return strings.Fields(s), nil
}

func main() {
	result, err := split("o joao roubou pao")
	fmt.Printf("result: %s err: %s", result, err)
}
