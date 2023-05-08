package main

import (
	"errors"
	"fmt"
	"sort"
)

// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com os valores ordenados de forma crescente. Caso o slice esteja vazio, retorne um erro.
func ordenarSlice(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice vazio")
	}
	novoSlice := make([]int, len(slice))
	copy(novoSlice, slice)
	sort.Ints(novoSlice)
	return novoSlice, nil
}
func main{
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	novoSlice, err := ordenarSlice(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(novoSlice)
	}
}
