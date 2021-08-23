package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int           { return len(x) }
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (x ordenarPorPotencia) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordenarPorNome []carro

func (x ordenarPorNome) Len() int           { return len(x) }
func (x ordenarPorNome) Less(i, j int) bool { return x[i].nome < x[j].nome }
func (x ordenarPorNome) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int           { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (x ordenarPorConsumo) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	carros := []carro{carro{"Chevete", 5000, 20},
		carro{"Porsche", 500, 80},
		carro{"Fusca", 3450, 10},
	}

	fmt.Println("Print sem ordenação:", carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Print ordenação por Potência:", carros)

	sort.Sort(ordenarPorNome(carros))

	fmt.Println("Print ordenação por Nome:", carros)

	sort.Sort(ordenarPorConsumo(carros))

	fmt.Println("Print ordenação por Consumo:", carros)

}
