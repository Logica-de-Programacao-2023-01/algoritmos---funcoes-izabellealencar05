package main

import (
	"errors"
	"strings"
)

// Escreva uma função que receba uma string como parâmetro e retorne uma nova string com todas as vogais substituídas por "*". Caso a string seja vazia, retorne um erro.
func replaceVowels(s string) (string, error) {
	if s == "" {
		return "", errors.New("string vazia")
	}

	vowels := "aeiouAEIOU"
	replacer := strings.NewReplacer(
		"a", "*", "e", "*", "i", "*", "o", "*", "u", "*",
		"A", "*", "E", "*", "I", "*", "O", "*", "U", "*",
	)

	return replacer.Replace(s), nil
}

func main() {
	s, err := replaceVowels("hello world")
	if err != nil {
		panic(err)
	}

	println(s) // h*ll* w*rld
}
