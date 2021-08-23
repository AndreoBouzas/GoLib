package main

import (
	"fmt"
	"sort"
)

//O package Sort Ordena uma Slice podendo ser aplicada todos os tipos primitivos de Go

func main() {

	//neste primeiro exemplo criamos uma slice de Strings

	ss := []string{"Andreo", "Diego", "Adriana", "Alex", "Mauro", "José", "Bianca", "Fernana", "Maria"}

	//por questão de controle, damos print na slice para demonstrar a "desordem"
	fmt.Println("Print antes do sort de Strings", ss)

	//para aplicar o sort, basta digitar o nome do pkg.nome_da_função(slice)
	sort.Strings(ss)

	//agora damos print do Slice após a função sort ter sido aplicada
	fmt.Println("Print do sort de Strings", ss)

	//////////////////////////////////////////////////////////////////////////////////////////////////
	/////////////////////////////SORT APLICADO A UMA SLICE DE INTS////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////

	//Neste segundo exemplo fazemos o mesmo com uma Slice de Ints

	//vamos criar uma slice de Ints
	sss := []int{1, 5, 7, 9, 4, 6, 3, 11, 4444, 2, 12, 88, 56, 52, 0, -1, 24}

	//por questão de controle, damos print na slice para demonstrar a "desordem"
	fmt.Println("Print antes do sort de Ints", sss)

	//para aplicar o sort, basta digitar o nome do pkg.nome_da_função(slice)
	sort.Ints(sss)

	//agora damos print do Slice após a função sort ter sido aplicada
	fmt.Println("Print do Sort de Ints", sss)

}
