package main

import "fmt"

func main() {

	var (
		soma3ou5 int
		soma3e5  int
		soma5    int
		soma3    int
		i        int
	)

	for i = 1; i < 1000; i++ {
		a := i % 3
		b := i % 5
		if a == 0 || b == 0 {
			soma3ou5 = soma3ou5 + i
			if a == 0 && b != 0 {
				soma3 = soma3 + i
			} else if a != 0 && b == 0 {
				soma5 = soma5 + i
			} else if a == 0 && b == 0 {
				soma3e5 = soma3e5 + i
			}

		}

	}
	fmt.Println("a soma de todos os múltiplos de 3 abaixo de", i, "é", soma3)
	fmt.Println("a soma de todos os múltiplos de 5 abaixo de", i, "é", soma5)
	fmt.Println("a soma de todos os múltiplos de 3 OU 5 abaixo de", i, "é", soma3ou5)
	fmt.Println("a soma de todos os múltiplos de 3 E 5 abaixo de", i, "é", soma3e5)

}
