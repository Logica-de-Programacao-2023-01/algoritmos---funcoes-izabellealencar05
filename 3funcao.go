package main

//Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.
import "fmt"

func concatStrings(s []string) string {
	var result string
	for _, str := range s {
		result += str
	}
	return result
}

func main() {
	s := []string{"Olá", " ", "mundo", "!"}
	fmt.Println(concatStrings(s)) // Output: "Olá mundo!"
}
