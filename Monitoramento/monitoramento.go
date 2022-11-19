package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const quantidadeMonitoramento = 3
const delay = 5

func main() {

	exibeIntruducao()
	//registraLog("site-falso", false)
	for {
		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimelogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0) // Para sair com sucesso.
		default:
			fmt.Println("Comando Inválido")
			os.Exit(-1) // quando deu alguma coisa errada
		}
	}
}

func exibeIntruducao() {
	nome := "Gabriel"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("A versão do programa de monitoramento é ", versao)
}

func lerComando() int {
	var novoComando int
	fmt.Scan(&novoComando)
	fmt.Println("A opção escolhida foi:", novoComando)

	return novoComando
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

// nessa funcao iniciamos o monitoramento, criamos um slice de string, fizemos um for(key, value) e mais abaixo
// fizemos outra funcao testeSites, para testar o sites pelo pacote http, usando o get(http.Get())

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := lerSites()
	//tem que fazer um for dentro do for para ele monitorar a quantidade de vezes que desejamos.
	for key := 0; key < quantidadeMonitoramento; key++ {
		for key, value := range site {
			fmt.Println("Testando site", key, ":", value)
			testeSites(value)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

// funcao para testar se o site está ativo, dando um get nos mesmos
func testeSites(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("O site está fora do ar", err)
	}

	if resp.StatusCode == 200 {
		fmt.Printf("O site: %s, foi carregado com sucesso\n", site)
		registraLog(site, true)
	} else {
		fmt.Printf("O site: %s, está com problema. Status code:%d\n", site, resp.StatusCode)
		registraLog(site, false)
	}
}

// arquivo, err := ioutil.ReadFile("sites1.txt") - poderiamos usar esse porém para o que precisamos o ideal e o pacote bufio
func lerSites() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Apresentou erro na função lerSites", err)
	}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Apresentou erro na função registra log", err)
	}

	// iremos utilizar o strconv para converter esse bool em string
	arquivo.WriteString(time.Now().Format("01-02-2006 15:04:05") + " - " + site + "- Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

	// nessa funcao os.OpenFile, ele recebe o nome do arquivo que vc quer abrir, as flags ( ou seja o que vc quer fazer)
	// os.O_RDWR - escrever e ler usamos os.O_RWDR | os.O_CREATE( | que significa ou)
	// e colocamos a permissao - 0666 - pode criar e ler
}

// Vai criar uma funcao, para abrir o arquivo, ler e imrpimir

func imprimelogs() {

	logs, err := ioutil.ReadFile("log.txt") // ler arquivo
	if err != nil {
		fmt.Println("Apresentou erro na função imprime logs", err)
	}
	fmt.Println(string(logs))

}
