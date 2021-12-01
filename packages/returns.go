package returns

import "fmt"

//definindo as variáveis de input do usuário
var num1, num2 int

//criando a função de calculo que recebe dois valores "int" e retorna 4 valores "int" e 4 valores "String"
func calc(x int, y int) (string, int, string, int, string, int, string, int) {
	return "\nSOMA =", x + y, "\nSUBTRAÇÃO =", x - y, "\nMULTIPLICAÇÃO =", x * y, "\nDIVISÃO =", x / y
}

//chamando na func main a função de multiplo retorno criada
func main() {
	//solicitando ao usuário que digite o primeiro número
	fmt.Println("Informe o primeiro número:")
	//coletado o primeiro número digitado
	fmt.Scanln(&num1)
	//solicitando ao usuário que digite o segundo número
	fmt.Println("Informe o Segundo número:")
	//coletado o segundo número digitado
	fmt.Scanln(&num2)
	//dando print do resultado da função calc com os valores informados pelo usuário
	fmt.Println(calc(num1, num2))
}
