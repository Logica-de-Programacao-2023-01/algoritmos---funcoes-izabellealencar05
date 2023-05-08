package main
//Escreva uma função que receba um slice de strings como parâmetro e retorne uma string com todas as strings concatenadas e separadas por vírgulas. Caso o slice esteja vazio, retorne um erro.
import (
	"errors"
	"fmt"
	"strings"
)

func joinStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("o slice está vazio")
	}
	return strings.Join(slice, ","), nil
}
func main{
	slice := []string{"abc", "def", "ghi"}
	result, err := joinStrings(slice)
	if err != nil {
		// tratamento do erro
	} else {
		// utilização do resultado
		fmt.Println(result)
	}
}

