package main

import "fmt"

func main() {

	//criando um canal que receberá os valores pares
	par := make(chan int)
	//criando um canal que receberá os valores impares
	impar := make(chan int)
	//criando um canal que receberá o sinal para interromper
	quit := make(chan bool)

	//chamando uma goroutine que envia valores pares, impares e um sinal quit
	go mandaNumeros(par, impar, quit)

	//chamando a função que recebe valores pares, impares e um sinal quit
	recive(par, impar, quit)

}

//definindo a função "mandaNumeros" que recebe um canal int de pares, impares e um canal de sinal quit
func mandaNumeros(par, impar chan int, quit chan bool) {
	//definindo uma variável de valor 100 para limitar o loop de valores enviados
	total := 100
	//criando o laço que contará até 100
	for i := 0; i < total; i++ {
		//definindo que caso o valor de "i" seja par o valor será enviado para o canal "par"
		if i%2 == 0 {
			par <- i
			//definindo que caso o valor de "i" seja impar o valor será enviado para o canal "impar"
		} else {
			impar <- i
		}
	}
	//após o final do laço fecha o canal par
	close(par)
	//após o final do laço fecha o canal impar
	close(impar)
	//após o final do laço envia "treu" para o canal quit
	quit <- true
}

//definindo a função recive que recebe um canal int de pares, impares e um canal de sinal quit
func recive(par, impar chan int, quit chan bool) {
	//definindo um loop "eterno" com select para definir a saída dos valores recebidos pelos canais
	for {
		select {
		//caso o valor tenha sido recebido pelo canal par, a mensagem a seguir será demosntrada.
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		//caso o valor tenha sido recebido pelo canal par, a mensagem a seguir será demosntrada.
		case v := <-impar:
			fmt.Println("O número", v, "é impar.")
		//caso o valor tenha sido recebido pelo canal quit
		case v, ok := <-quit:
			//verificamos se o "ok" recebido é falso e demonstramos uma mensagem correspondente
			if !ok {
				fmt.Println("Problema:", v)
				//senão a função encerra e deixa um aviso
			} else {
				fmt.Println("ENCERRANDO..... recebido:", v)
			}
			return
		}
	}
}
