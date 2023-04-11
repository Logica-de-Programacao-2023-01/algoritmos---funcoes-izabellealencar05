package main

import "fmt"

// Crie uma função que receba dois parâmetros (a e b) e retorne a divisão entre eles. Caso o
// segundo parâmetro seja zero, retorne um erro.
func main() {
	var a, b int
	fmt.Printf("escreva o valor de a: ")
	fmt.Scan(&a)
	fmt.Printf("escreva o valor de b: ")
	fmt.Scan(&b)

	resultado, err := divisao(a, b)
	if err != nil {
		fmt.Printf("ocorreu um erro ao dividir: %s\n", err)
		return
	}
	fmt.Printf("o resultado da divisao é: %d", resultado)

}
func divisao(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("nao pode dividir por 0")
	}
	return a / b, nil
}
