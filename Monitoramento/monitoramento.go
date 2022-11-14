package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Sejá bem-vindo ao programa de Monitoramento")
	fmt.Println("Versão 2.0")
	escolhaOpcao()

}

func escolhaOpcao() {

	var opcao string

	fmt.Println("Digite uma opção: \n1 - Para Monitorar.\n2 - Verificar Logs.\n3 - Sair do programa.")
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
