package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	Exibicao()

	for {
		escolhaOpcao()
	}
}

func escolhaOpcao() {

	var opcao string

	fmt.Println("Digite uma opção: \n1 - Para Monitorar.\n2 - Verificar Logs.\n0 - Sair do programa.")
	fmt.Scan(&opcao)

	switch opcao {
	case "1":
		fmt.Println("Monitorando...")

	case "2":
		fmt.Println("Verificando Logs...")
	case "0":
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida, encerrando programa")
		os.Exit(-1)
	}

}

func Exibicao() {
	name := "Gabriel Newton"
	fmt.Println("Sejam bem-vindo ao programa de Monitoramento")
	fmt.Println("Programa desenvolvido por", name, "- Versão 2.0")

}

func monitoramentoSites() {

}

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

func testaSites(site string) {

	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro na função testa Sites", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "apresentou erro no carregamento")
	}
}
