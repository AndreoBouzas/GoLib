package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Criando uma variável que recebe a função converge(recebe dois canais e devolve um canal)
	//A função trabalho recebe um string(mensagem) e devolve um canal2
	//Ao declararmos a variável canal ela esta usando a função converge que recebe a função trabahalho duas vezes, visto que a mesma responde um canal

	canal := converge(trabalho("One"), trabalho("Two"))

	//Crinando um laço que demonstra "n" vezes a variável canal.
	for x := 0; x < 16; x++ {

		fmt.Println(<-canal)
	}

}

//criando a função trabalho que recebe uma string de mensagem e devolve um canal string
func trabalho(s string) chan string {
	//criando um canal do tipo string
	canal := make(chan string)
	//lançando uma função como goroutine que recebe um string e um canal string
	go func(s string, c chan string) {
		//definindo um laço sem condição pra enviar para o canal "c" uma string externa junto do valor de "i"
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			//timer aleátorio entre 0 e 1 segundos para "exemplificar" o funcionamendo de um codeblock mais complexo
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
		//informado que o retorno da função é um srtring e um canal
	}(s, canal)

	return canal
}

//criando a função que recebe dois cainais e converge para um ao retornar um canal novo
func converge(x, y chan string) chan string {
	//criando uma canal novo
	novo := make(chan string)
	//crinado uma goroutine
	go func() {
		//definindo um laço infinto que coloca no novo canal tudo que sair do canal "x"
		for {
			novo <- <-x
		}
	}()
	//definindo um laço infinto que coloca no novo canal tudo que sair do canal "y"
	go func() {
		for {
			novo <- <-y
		}
	}()
	//retornando o novo canal
	return novo
}
