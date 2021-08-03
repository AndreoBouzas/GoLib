package main

import "fmt"

//criando a função de subtrair que recebe dois valores "int" e retorna um valor "int"
func sub(x int, y int) int {
	return x - y
}

//chamando a função sub e passando os parâmetros necessários.
func main() {
	fmt.Println(sub(66, 33))
}
