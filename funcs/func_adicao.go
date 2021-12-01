package main

import "fmt"

//criando a função de Somar que recebe dois valores "int" e retorna um valor "int"
func add(x int, y int) int {
	return x + y
}

//chamando a função add e passando os parâmetros necessários.
func main() {
	fmt.Println(add(66, 33))
}
