package main

import "fmt"

// Crie uma função que receba um número inteiro como parâmetro e retorne um novo slice contendo todos os números primos menores ou iguais a ele. Caso o número seja negativo, retorne um erro.

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeNumbers(n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("número inválido: %d", n)
	}
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes, nil
}

func main() {
	primes, err := primeNumbers(20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(primes)
}
