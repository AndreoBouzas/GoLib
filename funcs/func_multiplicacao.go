package main

import "fmt"

//criando a função de multiplicar que recebe dois valores "int" e retorna um valor "int"
func mult(x int, y int) int {
	return x * y
}

//chamando a função mult e passando os parâmetros necessários.
func main() {
	fmt.Println(mult(66, 33))
}
