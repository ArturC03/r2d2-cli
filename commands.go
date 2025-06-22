package main

import (
	"fmt"
	// "path/filepath"

	"github.com/ArturC03/r2d2"
	"github.com/ArturC03/r2d2Styles"
)

const Version = "0.1.1"

// ShowVersion displays the version of the R2D2 Language.
func ShowVersion() {
	fmt.Println(InfoMessage("R2D2 Language - version " + Version))
}

// Build simulates the compilation process.
func Build(r2d2Code string, filename string) (err error) {
	err = r2d2.BuildCode(r2d2Code, filename)

	if err != nil {
		return err
	}
	return nil
}

// Run simulates the execution process.
func Run(r2d2Code string) (err error) {
	err = r2d2.RunCode(r2d2Code)

	if err != nil {
		return err
	}
	return nil
}

// UnknownCommand shows an error message when an unknown command is entered.
func UnknownCommand(cmd string, errArgIndex int) {
	fmt.Println(("Unknown command: " + cmd))
	fmt.Println(HelpMessage("Run 'r2d2 help' for usage."))
}

// InfoMessage returns a styled info message.
func InfoMessage(message string) string {
	return r2d2Styles.InfoMessage(message)
}

// ErrorMessage returns a styled error message.
func ErrorMessage(message string) string {
	return r2d2Styles.ErrorMessage(message)
}

// HelpMessage returns a styled help message.
func HelpMessage(message string) string {
	return r2d2Styles.InfoMessage(message)
}

func BuildJs(input string, filename string) (err error) {
	err = r2d2.BuildJsFile(input, filename)

	if err != nil {
		return err
	}
	return nil
}
