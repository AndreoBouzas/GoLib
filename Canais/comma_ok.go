package main

import "fmt"

// A expressão comma ok serve para verificar se os valores recebidos no canal são verdadeiros, ou seja, se de fato foram recebidos
func main() {
	//criando um canal int
	canal := make(chan int)
	//lançando uma go func que envia o valor 0 para o canal
	go func() {
		canal <- 0
		//encerrando o canal após enviar um valor para ele
		close(canal)
	}()

	//criando duas variáveis que recebem guardam valores do canal. a variável "ok" é booleana e retorna false caso o valor recebido seja padrão do canal.
	v, ok := <-canal
	//dando print dos valores armazenados do canal
	fmt.Println(v, ok)
	v, ok = <-canal
	fmt.Println(v, ok)

	🤣 := "oi"
	fmt.Println(🤣)

	f:= unicode.Print(🤣)
	fmt.Println(f)
}
