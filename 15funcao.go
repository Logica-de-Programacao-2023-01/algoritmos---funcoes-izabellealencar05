package main

import (
	"errors"
	"fmt"
)

// Crie uma função que receba um número inteiro e um slice de inteiros como parâmetros e retorne verdadeiro se o número estiver presente no slice e falso caso contrário. Caso o slice esteja vazio, retorne um erro.
func findNumber(num int, nums []int) (bool, error) {
	if len(nums) == 0 {
		return false, errors.New("slice vazio")
	}

	for _, n := range nums {
		if n == num {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	nums := []int{2, 3, 5, 7, 11}
	num := 7

	found, err := findNumber(num, nums)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		if found {
			fmt.Printf("%d está presente no slice\n", num)
		} else {
			fmt.Printf("%d não está presente no slice\n", num)
		}
	}
}
