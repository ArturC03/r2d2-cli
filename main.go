// main.go
package main

import (
	"os"
	"strings"
)

// Remove this line as it's redeclared in commands.go
// const version = "0.0.1"

var commandLine = strings.Join(os.Args[1:], " ")

func main() {
    if len(os.Args) < 2 {
        ShowHelp()
        os.Exit(0)
    }

    cmd := os.Args[1]

    switch cmd {
    case "-help", "-h", "--help", "help", "--h":
        ShowHelp()
    case "-version", "-v", "--version", "version", "--v":
        ShowVersion()
    case "-build", "build":
        Build()
    case "-run", "run":
        Run()
    default:
        UnknownCommand(cmd, 1)
    }
}
