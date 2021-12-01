package main

import "fmt"

//criando a função de Dividir que recebe dois valores "int" e retorna um valor "int"
func div(x int, y int) int {

	return x / y
}

//chamando a função div e passando os parâmetros necessários.
func main() {
	fmt.Println(div(66, 33))
}
