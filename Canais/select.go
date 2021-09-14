package main

import "fmt"

func main() {
	//definindo x com valor de 500
	x := 500

	//definindo um canal "a" do tipo int
	a := make(chan int)
	//definindo um canal "b" do tipo int
	b := make(chan int)

	//definindo uma goroutine que mandará metade do valaor de x números
	go func(x int) {
		//definindo o que o laço enviará valores até que i seja do tamanho de x
		for i := 0; i < x; i++ {
			//enviando valores para o canal
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	//definindo o select onde ele indentifica de que canal está recebendo o valor e da print para informar
	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("canal A:", v)
		case v := <-b:
			fmt.Println("canal B:", v)
		}

	}
}
