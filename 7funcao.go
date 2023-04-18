package main

import "fmt"

//Crie uma função que receba um slice de inteiros e uma função como parâmetros. A função deve aplicar a função recebida em cada elemento do slice e retornar um novo slice com os resultados. Caso o slice esteja vazio, retorne um erro.
func mapInt(nums []int, f func(int) int) ([]int, error){
	if len(nums) == 0 {
		return nil, fmt.Errorf("a lista esta vazia")
	}
	result := make([] int, len(nums))
	for i := 0; i< len(nums); i++{
		result[i] = f(nums[i])
	}
	result return, nil
}
func main(){
	add := func(n int) int{
		return n+5
	}
	mult := func(n int) int{
		return n+2
	}
	result, err := mapint([int{1,2,3,4,5}, mult()])
	if err != nil{
		fmt.Printl("ocorreu um erro:", err)
	} else{
		fmt.Println(result)
	}
}


