package returns

import (
	"fmt"
)

var n1, n2 int
var sig string

func invoke() {

	fmt.Println("Informe o primeiro número:")
	fmt.Scanln(&n1)
	fmt.Println("Informe o operador (-,+,*,/ ) :")
	fmt.Scanln(&sig)
	fmt.Println("Informe o segundo número:")
	fmt.Scanln(&n2)

	fmt.Println("O resultado é:", (calculator(n1, n2, sig)))

	fmt.Println("O resultado das demais operaçõe é :", returns.calc(n1, n2))

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
