package main

import "errors"

// Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele for um número primo e falso caso contrário. Caso o número seja negativo, retorne um erro.
func isPrime(n int) (bool, error) {
	if n < 0 {
		return false, errors.New("número negativo")
	}

	if n < 2 {
		return false, nil
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}
