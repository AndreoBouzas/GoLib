package main

import "fmt"

func main() {

	//Criando um canal do tipo int
	c := make(chan int)

	//definindo a função meuloop como goroutine e passando os parâmetros necessários.
	go meuloop(10, c)

	//range fica esperando/ouvindo o canal e recebendo os valores passados
	for v := range c {
		fmt.Println("Recebido do canal:", v)
	}

}

//Criando uma função que deposita valores int em um canal
func meuloop(n int, s chan<- int) {
	//laço para enviar "n" valores para o canal
	for i := 0; i < n; i++ {
		//toda vez que o laço roda o canal recebe um valor
		s <- i
	}
	//Indica que não há mais nada a ser recebido e encerra o canal
	close(s)
}
