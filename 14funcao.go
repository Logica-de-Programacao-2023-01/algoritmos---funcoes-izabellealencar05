package main

import (
	"errors"
	"fmt"
)

func intersect(a []int, b []int) ([]int, error) {
	if len(a) == 0 || len(b) == 0 {
		return nil, errors.New("um dos slices está vazio")
	}

	m := make(map[int]bool)
	for _, num := range a {
		m[num] = true
	}

	var res []int
	for _, num := range b {
		if m[num] {
			res = append(res, num)
		}
	}

	return res, nil
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}

	res, err := intersect(a, b)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(res)
	}
}
