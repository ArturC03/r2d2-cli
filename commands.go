package main

import (
	"fmt"
)

const Version = "0.0.1"

// ShowVersion displays the version of the R2D2 Language.
func ShowVersion() {
	fmt.Println(InfoMessage("R2D2 Language - version " + Version))
}

// Build simulates the compilation process.
func Build() {
	fmt.Println(InfoMessage("Compiling..."))
}

// Run simulates the execution process.
func Run() {
	fmt.Println(InfoMessage("Executing..."))
}

// UnknownCommand shows an error message when an unknown command is entered.
func UnknownCommand(cmd string, errArgIndex int) {
	fmt.Println(ErrorMessage("Unknown command: " + cmd))
	fmt.Println(HelpMessage("Run 'r2d2 help' for usage."))
}

// InfoMessage returns a styled info message.
func InfoMessage(message string) string {
	return Tag("info") + " " + message
}

// ErrorMessage returns a styled error message.
func ErrorMessage(message string) string {
	return Tag("error") + " " + message
}

// HelpMessage returns a styled help message.
func HelpMessage(message string) string {
	return Tag("info") + " " + message
}
