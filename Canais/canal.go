package main

import "fmt"

func main() {

	//canais são a forma de comunicação entre duas goroutines

	//Um valor só poderá ser passado paro canal caso haja uma goroutine esperando receber/enviar algo pelo canal em questão

	//No exemplo abaixo uma goroutine esta passando o valor de 42 para o canal criado e a goroutine(main) esta recebendo o valor em questão

	//criando um canal de tipo int
	canal := make(chan int)
	//goroutine entregando um valor ao canal
	go func() { canal <- 42 }()
	//main (goroutine padrão)puxando o valor do canal
	fmt.Println(<-canal)

	//canias direcionais

	//Declarando um novo canal
	canaldirecional := make(chan int)

	//Chamando uma goroutine com canal criado
	go send(canaldirecional)

	//chamando  a função que está retirando informações do canal
	recive(canaldirecional)
}

//criando a função que deposita um valor int no canal
func send(s chan<- int) {
	//passando um valor pela variável "s" quereferencia o canal criado
	s <- 50
}

//criando uma função que retira um valor int no canal
func recive(r <-chan int) {
	//informa o valor retirado do canal
	fmt.Println("O valor recebido do canal foi:", <-r)

}
