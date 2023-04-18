package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Crie uma função que receba um número inteiro como parâmetro e retorne a soma de todos os seus dígitos. Caso o número seja negativo, retorne um erro.
func sumDigits(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("o numero é neativo")
	}
	s := strconv.Itoa(n)
	digits := strings.Split(s, "")
	for _, digit := range digits {
		d, err := strconv.Atoi(digit)
		if err != nil {
			return 0, err
		}
	}
}
