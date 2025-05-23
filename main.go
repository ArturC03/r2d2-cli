// main.go
package main

import (
	"fmt"
	"github.com/ArturC03/r2d2Styles"
	"os"
	"strings"
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

// parseOutputFlag checks for -o flag and returns the output filename
func parseOutputFlag() (string, bool) {
	for i, arg := range os.Args {
		if arg == "-o" && i+1 < len(os.Args) {
			return os.Args[i+1], true
		}
	}
	return "", false
}

// removeOutputFlag removes -o and its value from os.Args for compatibility
func removeOutputFlag() {
	var newArgs []string
	skip := false

	for i, arg := range os.Args {
		if skip {
			skip = false
			continue
		}
		if arg == "-o" && i+1 < len(os.Args) {
			skip = true // Skip the next argument (the output filename)
			continue
		}
		newArgs = append(newArgs, arg)
	}

	os.Args = newArgs
}

var commandLine = strings.Join(os.Args[1:], " ")

func main() {
	if len(os.Args) < 2 {
		ShowVersion()
		os.Exit(0)
	}

	// Parse -o flag before processing other arguments
	outputFile, hasOutput := parseOutputFlag()
	if hasOutput {
		removeOutputFlag()
	}

	cmd := os.Args[1]
	switch cmd {
	case "-help", "-h", "--help", "help", "--h":
		ShowHelp()
	case "-version", "-v", "--version", "version", "--v":
		ShowVersion()
	case "-b", "build":
		if hasOutput {
			Build(readR2D2File(), outputFile)
		} else {
			Build(readR2D2File(), getFilename())
		}
	case "-r", "run":
		Run(readR2D2File())
	case "js":
		if hasOutput {
			BuildJs(readR2D2File(), outputFile)
		} else {
			BuildJs(readR2D2File(), getFilename())
		}
	case "new":
		// MakeProject()
	default:
		UnknownCommand(cmd, 1)
	}
}
