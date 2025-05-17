// help.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Command represents a CLI command with its details
type Command struct {
	name        string
	description string
	usage       string
	examples    []string
	category    string
}

// Implementing list.Item interface
func (c Command) Title() string {
	return fmt.Sprintf("%s%s", highlightedCmdStyle.Render(c.name), categoryStyle.Render(" ["+c.category+"]"))
}
func (c Command) Description() string { return c.description }
func (c Command) FilterValue() string { return c.name + " " + c.description + " " + c.category }

// Command categories
const (
	CategoryBasic = "basic"
	CategoryBuild = "build"
	CategoryUtil  = "utility"
)

// Available commands
var commands = []Command{
	{
		"help",
		"Shows this interactive help menu",
		"r2d2 help",
		[]string{"r2d2 help --all"},
		CategoryBasic,
	},
	{
		"version",
		"Displays the language version",
		"r2d2 version",
		[]string{"r2d2 version --full", "r2d2 version -f"},
		CategoryBasic,
	},
	{
		"build",
		"Compiles a .r2d2 file",
		"r2d2 build <file.r2d2>",
		[]string{
			"r2d2 build hello.r2d2",
			"r2d2 build --optimize hello.r2d2",
		},
		CategoryBuild,
	},
	{
		"run",
		"Executes a .r2d2 file",
		"r2d2 run <file.r2d2>",
		[]string{
			"r2d2 run hello.r2d2",
			"r2d2 run --debug hello.r2d2",
		},
		CategoryBuild,
	},
	{
		"js",
		"Transpiles a .r2d2 file to JavaScript",
		"r2d2 js <file.r2d2>",
		[]string{
			"r2d2 js hello.r2d2",
			"r2d2 js --out hello.js hello.r2d2",
		},
		CategoryBuild,
	},
	{
		"init",
		"Initialize a new R2D2 project",
		"r2d2 init [project-name]",
		[]string{
			"r2d2 init my-project",
			"r2d2 init --template basic",
		},
		CategoryUtil,
	},
	// {
	// 	"package",
	// 	"Create a distributable package",
	// 	"r2d2 package [options]",
	// 	[]string{
	// 		"r2d2 package --target=all",
	// 		"r2d2 package --target=linux",
	// 	},
	// 	CategoryUtil,
	// },
}

// Styles for categories
var (
	categoryStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			Italic(true).
			MarginLeft(1)

	highlightedCmdStyle = lipgloss.NewStyle().
				Foreground(specialColor).
				Bold(true)

	exampleStyle = lipgloss.NewStyle().
			Foreground(infoColor).
			MarginLeft(2)
)

// Custom delegate for styling list items
type customDelegate struct {
	list.DefaultDelegate
}

func NewCustomDelegate() customDelegate {
	d := customDelegate{
		DefaultDelegate: list.NewDefaultDelegate(),
	}

	// Update styles using the colors from styles.go
	d.Styles.SelectedTitle = activeItemStyle
	d.Styles.NormalTitle = normalItemStyle
	d.Styles.DimmedTitle = normalItemStyle.Foreground(subtleColor)
	d.Styles.SelectedDesc = activeItemStyle.Bold(false)
	d.Styles.NormalDesc = normalItemStyle.Foreground(subtleColor)
	d.Styles.DimmedDesc = normalItemStyle.Foreground(subtleColor).Faint(true)

	return d
}

// HelpModel represents the application state
type HelpModel struct {
	list            list.Model
	help            string
	detailed        bool
	selectedItem    Command
	filterCategory  string
	categories      []string
	showingCategory bool
}

func (m HelpModel) Init() tea.Cmd {
	return nil
}

func (m HelpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle detailed view mode
	if m.detailed {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "q", "backspace":
				m.detailed = false
				m.help = "/ filter • c categories • q quit • ↑/↓ navigate • <CR> details"
				return m, nil
			}
		}
		return m, nil
	}

	// Handle category selection mode
	if m.showingCategory {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "c", "backspace":
				m.showingCategory = false
				return m, nil
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":
				idx := int(msg.Runes[0] - '1')
				if idx < len(m.categories) {
					// Filter by selected category
					m.filterByCategory(m.categories[idx])
					m.showingCategory = false
				}
				return m, nil
			case "0":
				// Clear category filter
				m.filterByCategory("")
				m.showingCategory = false
				return m, nil
			}
		}
		return m, nil
	}

	// Handle list view mode
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			if m.list.FilterState() == list.Filtering {
				m.list.ResetFilter()
				return m, nil
			}
			return m, tea.Quit
		case "/":
			m.help = "Type to filter • ESC to cancel • ENTER to select"
		case "c":
			m.showingCategory = true
			return m, nil
		case "enter":
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

func (m *HelpModel) filterByCategory(category string) {
	m.filterCategory = category

	// Create new filtered items list
	var filteredItems []list.Item
	if category == "" {
		// Show all commands
		filteredItems = make([]list.Item, len(commands))
		for i, cmd := range commands {
			filteredItems[i] = cmd
		}
	} else {
		// Show only commands matching the category
		for _, cmd := range commands {
			if cmd.category == category {
				filteredItems = append(filteredItems, cmd)
			}
		}
	}

	m.list.SetItems(filteredItems)
}

func (m HelpModel) View() string {
	// Category selector view
	if m.showingCategory {
		var sb strings.Builder
		sb.WriteString(headingStyle.Render("Select Category"))
		sb.WriteString("\n\n")

		for i, category := range m.categories {
			sb.WriteString(fmt.Sprintf("%s %s\n",
				highlightedCmdStyle.Render(fmt.Sprintf("%d.", i+1)),
				categoryStyle.Render(category),
			))
		}
		sb.WriteString(fmt.Sprintf("\n%s %s\n",
			highlightedCmdStyle.Render("0."),
			categoryStyle.Render("Show all commands"),
		))

		sb.WriteString("\n" + statusMessageStyle.Render("ESC to cancel"))

		return containerStyle.Width(50).Padding(1).Render(sb.String())
	}

	// Detailed view: show complete command information
	if m.detailed {
		detailBox := containerStyle.Width(65).Padding(1)

		title := lipgloss.NewStyle().
			Foreground(highlightColor).
			Bold(true).
			Render(m.selectedItem.name)

		category := lipgloss.NewStyle().
			Foreground(subtleColor).
			Italic(true).
			Render(m.selectedItem.category)

		description := lipgloss.NewStyle().
			Foreground(infoColor).
			Render(m.selectedItem.description)

		usage := lipgloss.NewStyle().
			Foreground(specialColor).
			Render(m.selectedItem.usage)

		// Examples section
		var examplesText string
		for _, example := range m.selectedItem.examples {
			examplesText += exampleStyle.Render(example) + "\n"
		}

		return detailBox.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				title,
				category,
				"",
				description,
				"",
				"Usage: "+usage,
				"",
				"Examples:",
				examplesText,
				statusMessageStyle.Render("ESC to return"),
			),
		)
	}

	// List view: show command list and help message
	return containerStyle.Width(40).Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.list.View(),
			statusMessageStyle.Render(m.help),
		),
	)
}

func ShowHelp() {
	// Get unique categories
	categoryMap := make(map[string]bool)
	for _, cmd := range commands {
		categoryMap[cmd.category] = true
	}

	categories := make([]string, 0, len(categoryMap))
	for category := range categoryMap {
		categories = append(categories, category)
	}

	// Convert commands to list.Item
	items := make([]list.Item, len(commands))
	for i, cmd := range commands {
		items[i] = cmd
	}

	// Setup the list model
	delegate := NewCustomDelegate()
	listModel := list.New(items, delegate, 30, 20)
	listModel.Title = "R2D2 Commands"
	listModel.SetShowPagination(false)
	listModel.Styles.Title = headingStyle
	listModel.SetFilteringEnabled(true)

	// Style the filter input
	listModel.FilterInput.PromptStyle = lipgloss.NewStyle().
		Foreground(highlightColor)
	listModel.FilterInput.TextStyle = lipgloss.NewStyle().
		Foreground(specialColor)
	listModel.FilterInput.Placeholder = "Filter..."

	// Initialize and run the program
	p := tea.NewProgram(
		HelpModel{
			list:           listModel,
			help:           "/ filter • c categories • q quit • ↑/↓ navigate • enter details",
			detailed:       false,
			categories:     categories,
			filterCategory: "",
		},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("%s %s\n", Tag("error"), err)
		os.Exit(1)
	}
}
