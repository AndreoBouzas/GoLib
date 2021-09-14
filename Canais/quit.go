package main

import "fmt"

func main() {
	//definindo um canal int bidirecional
	canal := make(chan int)
	//definindo um canal quit que recebrá um sinal de "finalização"
	quit := make(chan int)
	// definindo a função recebeQuit como uma goroutine que recebe um canal e um quit
	go recebeQuit(canal, quit)
	//chamando a função enviaPraCanal que enviará um canal e um quit
	enviaPraCanal(canal, quit)
}

// criando a função recebeQuit e definindo que ela recebe um canal do tipo int e um canal quit.
func recebeQuit(canal chan int, quit chan int) {
	//definindo um laço que irá esperar "x" valores
	for i := 0; i < 50; i++ {
		//print dos valores recebidos
		fmt.Println("recebido:", <-canal)
	}
	//ao final do laço é enviado um sinal de quit
	quit <- 1
}

//definindo uma função enviaPracanal e definindo que ela recebe um canal e um quit.
func enviaPraCanal(canal chan int, quit chan int) {
	//definindo uma variável qualquer para servir comocontador
	qualquercoisa := 0
	//criando um laço eterno
	for {
		//definindo um select
		select {
		// case que envia o valor da variável "qualquercoisa" para o canal
		case canal <- qualquercoisa:
			//acrecentando +1 a variável
			qualquercoisa++
			//case que verifica se foi recebido algo no canal quit
		case <-quit:
			return
		}
	}

}
