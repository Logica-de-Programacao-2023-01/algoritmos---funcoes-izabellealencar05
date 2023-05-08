package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Escreva uma função que receba um slice de strings como parâmetro e retorne uma string contendo apenas as strings que começam com uma letra maiúscula. Caso o slice esteja vazio, retorne um erro.
func filterUpperCase(strs []string) (string, error) {
	if len(strs) == 0 {
		return "", errors.New("slice vazio")
	}

	var result strings.Builder
	for _, str := range strs {
		if len(str) > 0 && unicode.IsUpper([]rune(str)[0]) {
			result.WriteString(str)
			result.WriteString(" ")
		}
	}

	return result.String(), nil
}

func main() {
	strs := []string{"ola", "meu", "nome", "é", "izabelle"}
	filtered, err := filterUpperCase(strs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(filtered)
}
