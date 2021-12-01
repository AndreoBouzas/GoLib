package main

import "fmt"

var n1, n2 int
var sig string

func main() {

	fmt.Println("Informe o primeiro número:")
	fmt.Scanln(&n1)
	fmt.Println("Informe o operador (-,+,*,/ ) :")
	fmt.Scanln(&sig)
	fmt.Println("Informe o segundo número:")
	fmt.Scanln(&n2)

	fmt.Println("O resultado é:", (calculator(n1, n2, sig)))

}

func calculator(num1, num2 int, operator string) int {
	result := 0
	if operator == "-" {
		result = num1 - num2
		return result
	} else if operator == "+" {
		result = num1 + num2
		return result
	} else if operator == "/" {
		result = num1 / num2
		return result
	} else if operator == "*" {
		result = num1 * num2
		return result
	} else {
		fmt.Println("Operador inválido")
	}
	return result
}
