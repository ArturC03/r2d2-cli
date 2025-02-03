// main.go
package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Command struct representa cada comando.
type Command struct {
	name        string
	description string
	usage       string
}

// Implementa a interface list.Item.
func (c Command) Title() string       { return c.name }
func (c Command) Description() string { return c.description }
func (c Command) FilterValue() string { return c.name + " " + c.description }

// Lista de comandos.
var commands = []Command{
	{"help", "Shows this help menu", "r2d2 help"},
	{"version", "Displays the language version", "r2d2 version"},
	{"build", "Compiles an .r2d2 file", "r2d2 build <file.r2d2>"},
	{"run", "Executes an .r2d2 file", "r2d2 run <file.r2d2>"},
}

// Delegate customizado para estilização.
type customDelegate struct {
	list.DefaultDelegate
}

func NewCustomDelegate() customDelegate {
	d := customDelegate{
		DefaultDelegate: list.NewDefaultDelegate(),
	}

	d.Styles.SelectedTitle = selectedItemStyle
	d.Styles.NormalTitle = itemStyle
	d.Styles.DimmedTitle = itemStyle.Copy().Foreground(lipgloss.Color("240"))
	d.Styles.SelectedDesc = selectedItemStyle.Copy().Bold(false)
	d.Styles.NormalDesc = itemStyle.Copy().Foreground(lipgloss.Color("238"))
	d.Styles.DimmedDesc = itemStyle.Copy().Foreground(lipgloss.Color("237"))

	return d
}

// HelpModel representa o estado da aplicação.
type HelpModel struct {
	list         list.Model
	help         string
	detailed     bool
	selectedItem Command
}

func (m HelpModel) Init() tea.Cmd {
	return nil
}

func (m HelpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Se estiver no modo detalhado, aguarda "esc" ou "q" para voltar.
	if m.detailed {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "q":
				m.detailed = false
				// Restaura a mensagem padrão.
				m.help = "/ to filter • q to quit • ↑/↓ to navigate • enter for usage"
				return m, nil
			}
		}
		return m, nil
	}

	// Modo de lista.
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			// Se estiver filtrando, reseta o filtro.
			if m.list.FilterState() == list.Filtering {
				m.list.ResetFilter()
				return m, nil
			}
			return m, tea.Quit
		case "/":
			m.help = "Type to filter • ESC to cancel • ENTER to select"
		case "enter":
			// Ao pressionar "enter", entra no modo detalhado.
			if i, ok := m.list.SelectedItem().(Command); ok {
				m.selectedItem = i
				m.detailed = true
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m HelpModel) View() string {
	// Modo detalhado: mostra os detalhes completos do item.

	if m.detailed {
		detailBox := lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(subtle).
			Padding(1, 2).
			Width(50)

		detail := lipgloss.NewStyle().
			Foreground(highlight).
			Bold(true).
			Render(fmt.Sprintf("Command: %s\n", m.selectedItem.name))

		description := lipgloss.NewStyle().
			Foreground(info).
			Bold(true).
			Render(fmt.Sprintf("Description: %s", m.selectedItem.description))

		usage := lipgloss.NewStyle().
			Foreground(special).
			Render("Usage: " + m.selectedItem.usage)

		instructions := messageStyle.Render("Press ESC or q to return")

		return detailBox.Render(
			lipgloss.JoinVertical(lipgloss.Left, detail, description, usage, instructions),
		)
	}
	// Modo de lista: mostra a lista de itens e a mensagem de ajuda.
	return baseStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.list.View(),
			messageStyle.Render(m.help),
		),
	)
}

func ShowHelp() {
	items := make([]list.Item, len(commands))
	for i, cmd := range commands {
		items[i] = cmd
	}

	delegate := NewCustomDelegate()
	listModel := list.New(items, delegate, 30, len(commands)*5)
	listModel.Title = "R2D2 CLI Commands"
	listModel.SetShowPagination(false)
	listModel.Styles.Title = titleStyle
	listModel.SetFilteringEnabled(true)

	// Personaliza os estilos do input de filtro.
	listModel.FilterInput.PromptStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("170"))
	listModel.FilterInput.TextStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252"))
	listModel.FilterInput.Placeholder = "Type / to filter commands..."

	p := tea.NewProgram(
		HelpModel{
			list:     listModel,
			help:     "/ to filter • q to quit • ↑/↓ to navigate • Enter for usage",
			detailed: false,
		},
		tea.WithAltScreen(),
	)

	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
