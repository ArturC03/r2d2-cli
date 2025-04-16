// main.go
package main

import (
	"fmt"
	"os"
	"strings"

	r2d2Styles "github.com/ArturC03/r2d2Styles"
)

// Remove this line as it's redeclared in commands.go
// const version = "0.0.1"

func readR2D2File() string {
	// Verifica se há pelo menos 3 argumentos na linha de comando (o primeiro é o nome do programa)
	if len(os.Args) < 3 {
		fmt.Println(r2d2Styles.ErrorMessage("Número insuficiente de argumentos"))
		fmt.Println(r2d2Styles.InfoMessage("Uso: programa <comando> <caminho_arquivo_r2d2>"))
		os.Exit(1)
	}

	// Obtém o caminho do arquivo (segundo argumento)
	filePath := os.Args[2]

	// Verifica se o caminho do arquivo está vazio
	if filePath == "" {
		fmt.Println(r2d2Styles.ErrorMessage("Caminho do arquivo não pode ser vazio"))
		os.Exit(1)
	}

	// Verifica se o arquivo existe
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Arquivo não encontrado - %s", filePath)))
		os.Exit(1)
	}

	// Lê o conteúdo do arquivo
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Erro ao ler o arquivo: %v", err)))
		os.Exit(1)
	}

	// Converte o conteúdo para string
	return string(content)
}

var commandLine = strings.Join(os.Args[1:], " ")

func main() {
	if len(os.Args) < 2 {
		ShowVersion()
		os.Exit(0)
	}

	cmd := os.Args[1]

	switch cmd {
	case "-help", "-h", "--help", "help", "--h":
		ShowHelp()
	case "-version", "-v", "--version", "version", "--v":
		ShowVersion()
	case "-b", "build":
		Build(readR2D2File())
	case "-r", "run":
		Run(readR2D2File())
	case "new":
		// MakeProject()
	default:
		UnknownCommand(cmd, 1)
	}
}
