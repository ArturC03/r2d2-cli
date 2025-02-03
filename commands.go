package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func ShowVersion() {
	fmt.Println(Tag("info"), "R2D2 Language - version 0.0.1")
}

func Build() { // TODO: Função que compila o código
	fmt.Println(Tag("info"), "Compiling...")
}

func Run() { // TODO: Função que executa o código
	fmt.Println(Tag("info"), "Executing...")
}

func UnknownCommand(cmd string, errArgIndex int) {
	errorMsg := lipgloss.NewStyle().
		Bold(true).
		Render("Unknown command: " + fmt.Sprint(lipgloss.NewStyle().Foreground(warning).Render(cmd)))

	helpMsg := lipgloss.NewStyle().
		Foreground(subtle).
		Render("Run 'r2d2 help' for usage.")

	fmt.Println(Tag("error"), errorMsg)
	fmt.Println(helpMsg)
}
