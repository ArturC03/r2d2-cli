// main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ArturC03/r2d2Styles"
)

// Gets the filename from the command line arguments
func getFilename() string {

	// Check if there are enough command line arguments
	if len(os.Args) < 3 {
		fmt.Println(r2d2Styles.ErrorMessage("Unsufficient number of arguments"))
		fmt.Println(r2d2Styles.InfoMessage("Args: r2d2 <comand> <file_path_r2d2> -o <output_file>"))
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

// Reads the contents of the R2D2 file
func readR2D2File() string {

	// Check if there are at least 3 command line arguments (the first is the program name)
	if len(os.Args) < 3 {
		fmt.Println(r2d2Styles.ErrorMessage("Isuficient number of arguments"))
		fmt.Println(r2d2Styles.InfoMessage("Use: r2d2 <comand> <file_path_r2d2>"))
		os.Exit(1)
	}

	// Gets the file path from the command line arguments
	filePath := os.Args[2]

	// Checks if the file path is empty
	if filePath == "" {
		fmt.Println(r2d2Styles.ErrorMessage("File path can't be empty"))
		os.Exit(1)
	}

	// Check if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("file not found: %v", filePath)))
		os.Exit(1)
	}

	// Reads the contents of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error reading file: %v", err)))
		os.Exit(1)
	}

	// Converts the contents to a string
	return string(content)
}

// Checks for -o flag and returns the output filename
func parseOutputFlag() (string, bool) {
	for i, arg := range os.Args {
		if arg == "-o" && i+1 < len(os.Args) {
			return os.Args[i+1], true
		}
	}
	return "", false
}

// Removes -o and its value from os.Args for compatibility
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

	var err error = nil
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
		// Check if there's a sub-argument for help
		if len(os.Args) > 2 && os.Args[2] == "static" {
			ShowHelpStatic()
		} else {
			ShowHelp()
		}

	case "-version", "-v", "--version", "version", "--v":
		ShowVersion()

	case "-b", "build":
		if hasOutput {
			err = Build(readR2D2File(), outputFile)
			if err != nil {
				os.Exit(1)
			}
			break
		}

		err = Build(readR2D2File(), getFilename())
		if err != nil {
			os.Exit(1)
		}
	case "-r", "run":
		err = Run(readR2D2File())
		if err != nil {
			os.Exit(1)
		}

	case "js":
		if hasOutput {
			err = BuildJs(readR2D2File(), outputFile)
			if err != nil {
				os.Exit(1)
			}
			break
		}

		err = BuildJs(readR2D2File(), getFilename())
		if err != nil {
			os.Exit(1)
		}
	case "new":
		// MakeProject() - je nes se'est pas
	default:
		UnknownCommand(cmd, 1)
	}
}
