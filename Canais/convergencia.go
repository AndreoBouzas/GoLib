package main

import (
	"fmt"
	"sync"
)

func main() {
	//Definindo um canal par
	par := make(chan int)
	//definindo um canal impar
	impar := make(chan int)
	//definindo um canal de convergeência que receberá de ambos canais
	converge := make(chan int)

	//chamando a função envia como uma goroutine que recebe um canal par e um impar
	go envia(par, impar)
	//chamando a função recebe como uma goroutine que recebe um canal par, impar e um canal de convergência
	go recebe(par, impar, converge)

	//Criando um laço que fica coletando valores do canal converge
	for v := range converge {
		//prit dos valores recebidos pelo canal
		fmt.Println("Valor recebido:", v)
	}

}

//criando a função envia que recebe um canal p e i
func envia(p, i chan int) {
	//definindo uma variável x de range máximo do laço.
	x := 100
	//ininciando o laço
	for n := 0; n < x; n++ {
		//verificando se a variável "n" é par
		if n%2 == 0 {
			//sendo par: "n" será enviado para o canal "p"
			p <- n
		} else {
			//sendo impar: "n" será enviado para o canal "i"
			i <- n
		}

	}
	//fechando os canais "p" e "i"
	close(p)
	close(i)

}

//criando a função recebe que converge os valores recebidos dos canais "p" e "i" para o canal "c"
func recebe(p, i, c chan int) {
	//criando uma waitGroup
	var wg sync.WaitGroup
	//informando o "lançamento" de uma goroutine
	wg.Add(1)
	//lançando um goroutine anônima para converger os valores do canal "p" para o canal "c"
	go func() {
		// crinado um laço range que fica "escutando" o canal "p" e envia qualquer resultado coletado para o canal "c"
		for v := range p {
			c <- v
		}
		//informando o fim da goroutine "lançada"
		wg.Done()
	}()
	//informando o "lançamento" de uma goroutine
	wg.Add(1)
	//lançando um goroutine anônima para converger os valores do canal "i" para o canal "c"
	go func() {
		// crinado um laço range que fica "escutando" o canal "i" e envia qualquer resultado coletado para o canal "c"
		for v := range i {
			c <- v
		}
		//informando o fim da goroutine "lançada"
		wg.Done()
	}()
	//informando que é necessário a conclusão de todas as goroutines lançadas para prosseguir
	wg.Wait()
	//fechando o canal "c"
	close(c)
}
