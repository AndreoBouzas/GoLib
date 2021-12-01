package main

import "fmt"

//criando a função de Somar que recebe dois valores "int" e retorna um valor "int"
func add(x int, y int) int {
	return x + y
}

//criando a função de subtrair que recebe dois valores "int" e retorna um valor "int"
func sub(x int, y int) int {
	return x - y
}

//criando a função de multiplicar que recebe dois valores "int" e retorna um valor "int"
func mult(x int, y int) int {
	return x * y
}

//criando a função de Dividir que recebe dois valores "int" e retorna um valor "int"
func div(x int, y int) int {

	return x / y
}

//chamando na func main as funções criadas  para cada caso
func main() {
	fmt.Println(add(66, 33))
	fmt.Println(sub(66, 33))
	fmt.Println(mult(66, 33))
	fmt.Println(div(66, 33))
}
