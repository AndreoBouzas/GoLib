package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"

	"log"
	"os"
	"time"
)

//Criando as variáveis necessárias
var (
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	//grav
)

func main() {
	//Criando um arquivo com o nome teste.txt > se não for definido um caminho o arquivo será criado no mesmo diretório do código
	newFile, err = os.Create("test.txt")
	//verificando se ocorreu um erro na criação do arquivo
	if err != nil {
		//caso sim encerra o código informando o erro
		fmt.Println(err)
	}

	//Criando uma variável que receberá os status do arquivo nomeado teste.txt
	fileInfo, err = os.Stat("test.txt")
	//verificando se ocorreu um erro na criação do arquivo
	if err != nil {
		//caso sim encerra o código informando o erro
		fmt.Println(err)
	}

	log.Println(newFile)

	//Print do método name do pkg os.FileInfo (demonstra o nome do arquivo)***** sintaxe > variável_que_contem_o_os.Stat_do_arquivo.Name()
	fmt.Println("Nome do Arquivo:", fileInfo.Name())
	//Print do método Size do pkg os.FileInfo (demonstra o Tamanho do arquivo)***** sintaxe > variável_que_contem_o_os.Stat_do_arquivo.Size()
	fmt.Println("Tamanho em Bytes do arquivo:", fileInfo.Size())
	//Print do método Mode do pkg os.FileInfo (demonstra as permissões do arquivo)***** sintaxe > variável_que_contem_o_os.Stat_do_arquivo.Mode()
	fmt.Println("Permissões:", fileInfo.Mode())
	//Print do método ModTime do pkg os.FileInfo (verifica a última modificação feita no rquivo)***** sintaxe > variável_que_contem_o_os.Stat_do_arquivo.ModTime()
	fmt.Println("Última modificação:", fileInfo.ModTime())
	//Print do método IsDir do pkg os.FileInfo (Verifica se o arquivo é um diretório)***** sintaxe > variável_que_contem_o_os.Stat_do_arquivo.IsDir()
	fmt.Println("É diretório: ", fileInfo.IsDir())

	//Fechando o arquivo
	newFile.Close()

	//---------------------------------------------------------------------------------------------------------------------//
	//////////////////////////////////////////Leitura simples de um arquivo!/////////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	file, err := os.Open("test.txt") // SINTAXE > vairável = os.Open("nome_do_arquivo.txt")

	//verificando se ocorreu um erro na criação do arquivo
	if err != nil {
		//caso sim encerra o código informando o erro
		fmt.Println(err)
	}
	//fechando o arquivo
	file.Close()

	//Declaração de leitura do arquivo com mais opções
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666) //SINTAXE > vairável = os.OpenFile("nome_do_arquivo.txt", parâmetros, modo de permissão)

	// É possível utilizar um ou mais parâmetros, basta adicionar o caracter "|"
	// Ex > os.O_CREATE|os.O_APPEND
	// Ex > os.O_CREATE|os.O_TRUNC|os.O_WRONLY

	//Parâmetros

	// os.O_CREATE // Cria o arquivo caso o mesmo não exista!
	// os.O_RDONLY // Apenas Leitura
	// os.O_WRONLY // Apenas Escrita
	// os.O_RDWR // Ler e Escrever
	// os.O_APPEND // concatena ao fim do arquivo
	// os.O_TRUNC // Truncar arquivo ao abrir

	//verificando se ocorreu um erro na criação do arquivo
	if err != nil {
		//caso sim encerra o código informando o erro
		fmt.Println(err)
	}
	//fechando o arquivo
	file.Close()

	// Abrindo o arquivos com o parâmetro de escrita.
	file, err = os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		// Testando se é possível escrever no arquivo.
		if os.IsPermission(err) {
			//Informando o erro encontrado ao verificar a permissão para escrever
			log.Println("Erro: Permissão de Escrita negada!.")
		}
	}
	//fechando o arquivo
	file.Close()

	//---------------------------------------------------------------------------------------------------------------------//
	///////////////////////////////////Verificando as permissões de um arquivo!//////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	//Se o mesmo não existir o erro retornado será diferente!!
	//Para estes casos é ideal utilizar o parâmetro "os.IsNotExist(err)" para verificar isto

	// Abrindo o arquivos com o parâmetro de leitura.
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		// Testando se é possível Ler o arquivo.
		if os.IsPermission(err) {
			//Informando o erro encontrado ao verificar a permissão para ler
			log.Println("Erro: Permissão de Leitura negada!")
		}
	}
	//fechando o arquivo
	file.Close()

	//---------------------------------------------------------------------------------------------------------------------//
	//////////////////////////////////Como alterar as permissões de um arquivo!//////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	//É possível alterar as permissões de um arquivo usando a Sintaxe Linux!
	//Para isto vamos utilizar o comando "os.Chmod("Nome_Do_Arquivo.txt", 0777"padrão de permissão linux")"

	err = os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	//Trocando a informação de data de mofificação do arquivo
	doisDiasAPartirDeAgora := time.Now().Add(48 * time.Hour)
	momentoDoUltimoAcesso := doisDiasAPartirDeAgora
	momentoDaUltimaModificacao := doisDiasAPartirDeAgora
	err = os.Chtimes("test.txt", momentoDoUltimoAcesso, momentoDaUltimaModificacao)
	if err != nil {
		log.Println(err)
	}

	//---------------------------------------------------------------------------------------------------------------------//
	//////////////////////////////////Buscando posições em um arquivo!///////////////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	//Inicialmente devamos abrir o arquivo em questão usando o os.Open("Nome_Do_Arquivo.txt")
	file, _ = os.Open("test.txt")

	//aqui definimos que o arquivo será fechado ao fim do programa
	defer file.Close()

	// Á variável "deslocamento" representa quantos bytes mover (Podenso ser representado por um valor positivo ou negativo)
	var deslocamento int64 = 5

	// A variável "deonde" é o ponto de referência para a variável "deslocamento"
	// 0 = Indica o Início do arquivo
	// 1 = Indica a posição atual de leitura do arquivo
	// 2 = Indica o Fim do arquivo

	var deonde int = 0

	novaPosicao, err := file.Seek(deslocamento, deonde)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Moveu-se apenas 5:", novaPosicao)

	// voltando 2 bytes da posição atual
	novaPosicao, err = file.Seek(-2, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Voltou 2 bytes:", novaPosicao)

	// Encontrando a posição atual através do resultado "return" do valor de "Seek" após mover 0 bytes
	posicaoAtual, err := file.Seek(0, 1)
	fmt.Println("Posição Atual:", posicaoAtual)

	// Voltando para o início do arquivo
	novaPosicao, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Posição Atual após voltar para o ínicio do arquivo", novaPosicao)

	//---------------------------------------------------------------------------------------------------------------------//
	//////////////////////////////////Como Gravar Bytes em um arquivo!///////////////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	// Abrindo um novo arquivo com permissão de apenas escrita
	filee, erro := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(erro)
	}
	defer file.Close()

	//Escrevendo os bytes no arquivo
	byteSlice := []byte("Bytes!\n")
	bytesEscritos, err := filee.Write(byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Escrito %d bytes.\n", bytesEscritos)

	//O pacote ioutil tem uma função útil chamada "WriteFile ()"" que tratará da criação / abertura, gravação de uma Slice de bytes e fechamento.
	// É útil se você só precisa de uma maneira rápida de despejar uma Slice de bytes em um arquivo.

	err = ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		fmt.Println(err)
	}

	//---------------------------------------------------------------------------------------------------------------------//
	/////////////////////////////////////////Usar Gravador com Buffer!///////////////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	//O pacote bufio permite criar um gravador com buffer para que você possa trabalhar com um buffer na memória antes de gravá-lo no disco.
	//Isso é útil se você precisar fazer muita manipulação nos dados antes de gravá-los no disco para economizar tempo do disco I/O.
	// Também é útil se você gravar apenas um byte por vez e quiser armazenar um grande número na memória antes de despejá-lo no arquivo de uma vez,
	//caso contrário, você executaria I/O de disco para cada byte. Isso causa desgaste no disco e também retarda o processo.

	// Abrindo arquivo para leitura
	file, err = os.OpenFile("test.txt", os.O_WRONLY, 0666)

	//Testando possível erro durante a leitura do arquivo
	if err != nil {
		fmt.Println(err)
	}

	//definindo que o arquivo será fechado somente quando o programa encerrar
	defer file.Close()

	// Criando um gravador em buffer a partir do arquivo
	bufferedWriter := bufio.NewWriter(file)

	// Grava bytes no buffer
	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Bytes Escritos: %d\n", bytesWritten)

	//---------------------------------------------------------------------------------------------------------------------//
	/////////////////////////////////////////Usar Gravador de String com Buffer!///////////////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	// Escrevendo uma String para o buffer
	bytesWritten, err = bufferedWriter.WriteString(
		"String em buffer\n",
	)

	//Verificando possível erro
	if err != nil {
		fmt.Println(err)
	}

	//Print dos Bytes escritos
	log.Printf("Bytes escritos: %d\n", bytesWritten)

	// Verificando quanto está armazenado no buffer em espera
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes Armazenados : %d\n", unflushedBufferSize)

	// SVerificando quanto buffer está disponível
	bytesAvailable := bufferedWriter.Available()

	//Verificando possível erro
	if err != nil {
		fmt.Println(err)
	}
	//Print do buffer disponível
	log.Printf("buffer Disponível: %d\n", bytesAvailable)

	// Grava o buffer de memória no disco
	bufferedWriter.Flush()

	// Reverta quaisquer alterações feitas no buffer que ainda não foram gravadas no arquivo com Flush ().
	// Acabamos de liberar, portanto, não há mudanças para reverter
	// O gravador que você passa como um argumento é para onde o buffer fará a saída, se você quiser mudar para um novo gravador, use:
	bufferedWriter.Reset(bufferedWriter)

	// Veja quanto buffer está disponível
	bytesAvailable = bufferedWriter.Available()

	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("buffer Disponível: %d\n", bytesAvailable)

	// Redimencionar o buffer.
	//O primeiro argumento é um gravador para onde o buffer deve enviar a saída.
	//Neste caso, estamos usando o mesmo buffer.
	//Se escolhermos um número menor do que o buffer existente, como 10, não obteremos um buffer de tamanho 10.
	//Obteremos de volta um buffer do tamanho do original, pois já é grande o suficiente (padrão 4096)
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)

	// Verificando o tamanho do buffer disponível após o redimensionamento
	bytesAvailable = bufferedWriter.Available()

	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("buffer Disponível após o redimensionamento:  %d\n", bytesAvailable)

	//---------------------------------------------------------------------------------------------------------------------//
	/////////////////////////////////////////copiar de um arquivo para o outro //////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	//para isto devemos fazer o seguinte
	// abrir o arquivo original usando o comando simples "os.Open(Nome_Do_Arquivo.txt)"
	originalFile, err := os.Open("test.txt")
	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}

	//Aqui declaramos que o arquivo "OriginalFile será fechado somente após o término do programa"
	defer originalFile.Close()

	//criando um novo arquivo
	newFile, err := os.Create("test_copy.txt")
	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}
	//Aqui declaramos que o arquivo "newFile será fechado somente após o término do programa"
	defer newFile.Close()

	// Copiando os bytes da fonte para o destino, neste caso do arquivo "newFile" para o arquivo "originalFile"
	//note a sintaxe usada no comando "os.Copy(Arquivo_De_Destino, Arquivo_Fonte)"
	bytes_Escritos, err := io.Copy(newFile, originalFile)
	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Copiou %d bytes.", bytes_Escritos)

	// A função "os.Sync" serve para confirmar o conteúdo atual do arquivo para armazenamento estável.
	// Normalmente, isso significa liberar a cópia na memória do sistema de arquivos de dados gravados recentemente no disco.
	err = newFile.Sync()
	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}

	//---------------------------------------------------------------------------------------------------------------------//
	/////////////////////////////////////////Lendo bytes específicos do ARQUIVO /////////////////////////////////////////////
	//---------------------------------------------------------------------------------------------------------------------//
	//---------------------------------------------------------------------------------------------------------------------//

	//O tipo os.File fornece algumas funções básicas. Os pacotes io , ioutil e bufio fornecem funções adicionais para trabalhar com arquivos.

	// Abrindo o arquivo paraa leitura do mesmo
	file, err = os.Open("test.txt")
	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	////////////////////////////////////// Lendo até o len(b) bytes do arquivo////////////////////////////////////////////
	// Zero bytes Escritos significa o fim do arquivo
	// O fim do arquivo retorna o tipo de erro io.EOF

	//definindo uma variável slice do tipo byte com tamanho de 16
	fatia_De_Bytes := make([]byte, 16)

	bytesRead, err := file.Read(fatia_De_Bytes)
	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Número de bytes lidos: %d\n", bytesRead)
	log.Printf("Informação lida: %s\n", fatia_De_Bytes)

	////////////////////////////////////// Lendo exatamente "n" bytes do arquivo////////////////////////////////////////////

	// Abrindo o arquivo para leitura
	file, err = os.OpenFile("test1.txt", os.O_CREATE, 0777)

	//Tratando possível erro na abertura do arquivo
	if err != nil {
		fmt.Println(err)
	}

	// A função file.Read () lerá um arquivo minúsculo em uma grande fatia de bytes.
	// mas io.ReadFull () retornará um erro se o arquivo for menor que a fatia de bytes.

	//definindo uma variável slice do tipo byte com tamanho de 2
	fatia_de_byte := make([]byte, 2)

	num_De_Bytes_Lidos, err := io.ReadFull(file, fatia_de_byte)
	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Número de bytes lidos: %d\n", num_De_Bytes_Lidos)
	log.Printf("Informação lida: %s\n", fatia_de_byte)

	////////////////////////////////////// Lendo pelo menos "n" bytes do arquivo////////////////////////////////////////////

	// Abrindo o arquivo para ler
	arquivo_de_teste, err := os.Open("sketch.txt")

	//Tratando possível erro na abertura do arquivo
	if err != nil {
		fmt.Println(err)
	}

	//definindo uma variável slice do tipo byte com tamanho de 2
	fatia_de_Byte := make([]byte, 512)

	//definindo um valor minimo de bytes
	minBytes := 8

	//io.ReadAtLeast () retornará um erro se não puder encontrar pelo menos minBytes para ler.
	//Ele lerá tantos bytes quanto a fatia_de_Byte puder conter.
	numBytesRead, err := io.ReadAtLeast(arquivo_de_teste, fatia_de_Byte, minBytes)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Número de bytes lidos: %d\n", numBytesRead)
	log.Printf("Informação lida: %s\n", fatia_de_Byte)

	////////////////////////////////////// Lendo TODOS os bytes de um arquivo////////////////////////////////////////////

	// Abrindo o Arquivo para leitura
	arquivo_Com_Todos_Bytes, err := os.Open("sketch.txt")

	//Tratando possível erro na abertura do arquivo
	if err != nil {
		fmt.Println(err)
	}

	// os.File.Read(), io.ReadFull(), e io.ReadAtLeast() funcionam com uma Slice de bytes fixa que você cria antes de ler o arquivo

	// ioutil.ReadAll() lerá cada byte do leitor (neste caso, um arquivo) e retornará uma Slice de Slice desconhecida
	data, err := ioutil.ReadAll(arquivo_Com_Todos_Bytes)
	//Tratando possível erro
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Informação em hexadecimal: %x\n", data)
	fmt.Printf("Informação em String: %s\n", data)
	fmt.Println("Número de bytes lidos:", len(data))

}
