package main

import "fmt"

//criando uma Struct Pessoa
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

//Criando uma struct Dentista que contem pessoa
type dentista struct {
	pessoa
	especializado string
	salario       float64
}

//Criando uma struct arquiteto que contem pessoa
type arquiteto struct {
	pessoa
	estilodeconstrucao string
	orcamento          float64
}

//Criando um método para struct dentista
func (x dentista) apresentacao() {
	fmt.Println("Meu nome é", x.nome, " e eu sou especialista em ", x.especializado, ", Tenha um bom dia!")
}

//criando um método para stuct arquiteto
func (x arquiteto) apresentacao() {
	fmt.Println("Meu nome é", x.nome, "e meu orçamento é ", x.orcamento, ", Tenha um bom dia!")

}

//criando outro método para stuct arquiteto
func (x arquiteto) apresentacao2() {
	fmt.Println("entendo sua frustração, más para construir algo baseada na arquitetura", x.estilodeconstrucao, "de fato sai o preço já informado")
}

//criando a interface "implementacao1" que define que tudo que possuir o método "apresentacao" e "apresentacao2" será implementado
type implementacao1 interface {
	apresentacao()
	apresentacao2()
}

//criando a interface "implementacao2" que define que tudo que possuir o método "apresentacao" será implementado
type implementacao2 interface {
	apresentacao()
}

//criando a interface que define que tudo que possuir os seguintes métodos são deste "tipo"
func amigo(m implementacao2) {
	m.apresentacao()

}

func profissional(m implementacao1) {
	m.apresentacao()
	m.apresentacao2()
}

func main() {

	//populando a struct arquiteto
	arquiteto1 := arquiteto{
		pessoa: pessoa{
			nome:      "Antônio",
			sobrenome: "Andrade",
			idade:     21,
		},
		estilodeconstrucao: "industruial moderno",
		orcamento:          9510.68,
	}

	//populando a struct dentista
	dentista1 := dentista{
		pessoa: pessoa{
			nome:      "Rogério",
			sobrenome: "Cardoso",
			idade:     65,
		},
		especializado: "tratamento de canal",
		salario:       2700.68,
	}

	//chamando arquiteto pelo interface "amigo"
	amigo(arquiteto1)
	//chamando arquiteto pelo interface "profissional"
	profissional(arquiteto1)

	//como o método "apresentacao" é implementado por dois intarfaces os structs que forem utilizados por eles serão correspondidos

	//chamando dentista pelo método e não pela interface
	dentista1.apresentacao()
	//arquiteto1.apresentacao()
}







		
			





