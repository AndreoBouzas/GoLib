package main

import "fmt"

//criando a função de calculo que recebe dois valores "int" e retorna 4 valores "int" e 4 valores "String"
func calc(x int, y int) (string, int, string, int, string, int, string, int) {
	return "\nSOMA =", x + y, "\nSUBTRAÇÃO =", x - y, "\nMULTIPLICAÇÃO =", x * y, "\nDIVISÃO =", x / y
}

//chamando na func main a função de multiplo retorno criada
func main() {
	fmt.Println(calc(66, 33))
}
