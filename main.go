package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	version = "0.0.1"
)

var (
	commandLine string = strings.Join(os.Args[1:], " ")
)

func main() {
	fmt.Println()
	if len(os.Args) < 2 {
		ShowHelp()
		os.Exit(0)
	}

	args := os.Args[1:]

	switch args[0] {
	case "-help", "-h", "--help", "help", "--h":
		ShowHelp()
	case "-version", "-v", "--version", "version", "--v":
		ShowVersion()
	case "-build", "build":
		Build()
	case "-run", "run":
		Run()
	default:
		UnknownCommand(args[0], 0+1)
	}
}
