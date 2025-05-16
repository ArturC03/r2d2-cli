// main.go
package main

import (
	"fmt"
	"os"
	"strings"

	r2d2Styles "github.com/ArturC03/r2d2Styles"
)

func getFilename() string {
	// Check if there are enough command line arguments
	if len(os.Args) < 3 {
		fmt.Println(r2d2Styles.ErrorMessage("Número insuficiente de argumentos"))
		fmt.Println(r2d2Styles.InfoMessage("Args: program <comand> <file_path_r2d2>"))
		os.Exit(1)
	}

	// Get the file path (second argument)
	filePath := os.Args[2]

	// Check if the file path is empty
	if filePath == "" {
		fmt.Println(r2d2Styles.ErrorMessage("File path can't be empty"))
		os.Exit(1)
	}

	// Extract just the filename from the path
	parts := strings.Split(filePath, string(os.PathSeparator))
	filename := parts[len(parts)-1]

	return filename
}

func readR2D2File() string {
	// Verifica se há pelo menos 3 argumentos na linha de comando (o primeiro é o nome do programa)
	if len(os.Args) < 3 {
		fmt.Println(r2d2Styles.ErrorMessage("Isuficient number of arguments"))
		fmt.Println(r2d2Styles.InfoMessage("Use: program <comand> <file_path_r2d2>"))
		os.Exit(1)
	}

	// Obtém o caminho do arquivo (segundo argumento)
	filePath := os.Args[2]

	// Verifica se o caminho do arquivo está vazio
	if filePath == "" {
		fmt.Println(r2d2Styles.ErrorMessage("File path can't be empty"))
		os.Exit(1)
	}

	// Verifica se o arquivo existe
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("file not found: %v", filePath)))
		os.Exit(1)
	}

	// Lê o conteúdo do arquivo
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error reading file: %v", err)))
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
		Build(readR2D2File(), getFilename())
	case "-r", "run":
		Run(readR2D2File())
	case "js", "buildjs":
		BuildJs(readR2D2File(), getFilename())
	case "new":
		// MakeProject()
	default:
		UnknownCommand(cmd, 1)
	}
}
