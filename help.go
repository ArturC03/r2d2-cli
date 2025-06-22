package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ArturC03/r2d2Styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// Represents a CLI command with its details
type Command struct {
	name        string
	description string
	usage       string
	examples    []string
	category    string
}

// Implementing list.Item interface - return plain text for filtering
func (c Command) Title() string {
	return fmt.Sprintf("%s [%s]", c.name, c.category)
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
		"Shows help menu (interactive by default, use 'static' for simple output)",
		"r2d2 help [static]",
		[]string{"r2d2 help", "r2d2 help static"},
		CategoryBasic,
	},
	{
		"version",
		"Displays the language version",
		"r2d2 version",
		[]string{
			"r2d2 version",
			// "r2d2 version -f"
		},
		CategoryBasic,
	},
	{
		"build",
		"Compiles a .r2d2 file",
		"r2d2 build <file.r2d2>",
		[]string{
			"r2d2 build hello.r2d2",
			"r2d2 build hello.r2d2 -o hi",
			// "r2d2 build --optimize hello.r2d2",
		},
		CategoryBuild,
	},
	{
		"run",
		"Executes a .r2d2 file",
		"r2d2 run <file.r2d2>",
		[]string{
			"r2d2 run hello.r2d2",
			// "r2d2 run debug hello.r2d2",
		},
		CategoryBuild,
	},
	{
		"js",
		"Transpiles a .r2d2 file to JavaScript",
		"r2d2 js <file.r2d2>",
		[]string{
			"r2d2 js hello.r2d2",
			"r2d2 js hello.r2d2 -o bye.js",
		},
		CategoryBuild,
	},
	// {
	// 	"init",
	// 	"Initialize a new R2D2 project",
	// 	"r2d2 init [project-name]",
	// 	[]string{
	// 		"r2d2 init my-project",
	// 		"r2d2 init --template basic",
	// 	},
	// 	CategoryUtil,
	// },
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

// Styles for categories (using colors from styles.go)
var (
	categoryStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			Italic(true)

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
	width int
}

func NewCustomDelegate(width int) customDelegate {
	d := customDelegate{
		DefaultDelegate: list.NewDefaultDelegate(),
		width:           width,
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

// Custom render method to handle styling properly
func (d customDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var str string

	if cmd, ok := item.(Command); ok {
		// Apply custom styling here instead of in Title()
		cmdName := highlightedCmdStyle.Render(cmd.name)

		// Get current terminal width from the model width (we'll estimate it)
		termWidth := 80 // default fallback
		if m.Width() > 0 {
			termWidth = m.Width()
		}

		// Responsive category display
		var category string
		if termWidth < 50 {
			// Very compact for small screens
			category = categoryStyle.Render(" [" + string(cmd.category[0]) + "]")
		} else if termWidth < 70 {
			// Short category names for medium screens
			shortCategory := cmd.category
			if len(shortCategory) > 5 {
				shortCategory = shortCategory[:5]
			}
			category = categoryStyle.Render(" [" + shortCategory + "]")
		} else {
			// Full category for larger screens
			category = categoryStyle.Render(" [" + cmd.category + "]")
		}

		// Responsive description handling
		description := cmd.description
		if termWidth < 60 && len(description) > 40 {
			description = description[:37] + "..."
		} else if termWidth < 80 && len(description) > 60 {
			description = description[:57] + "..."
		}

		if index == m.Index() {
			// Selected item
			str = d.Styles.SelectedTitle.Render(cmdName + category)
			if description != "" && termWidth > 40 {
				str += "\n" + d.Styles.SelectedDesc.Render(description)
			}
		} else {
			// Normal item
			str = d.Styles.NormalTitle.Render(cmdName + category)
			if description != "" && termWidth > 40 {
				str += "\n" + d.Styles.NormalDesc.Render(description)
			}
		}
	} else {
		// Fallback to default rendering
		d.DefaultDelegate.Render(w, m, index, item)
		return
	}

	fmt.Fprint(w, str)
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
	width           int
	height          int
}

func (m HelpModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m HelpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle window size changes
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Update list dimensions dynamically
		listWidth := m.width - 4
		if listWidth < 30 {
			listWidth = 30
		}
		if listWidth > 80 {
			listWidth = 80
		}

		listHeight := m.height - 8
		if listHeight < 10 {
			listHeight = 10
		}
		if listHeight > 25 {
			listHeight = 25
		}

		m.list.SetSize(listWidth, listHeight)

		// Note: Delegate width will be updated on next render

		// Update help message based on screen width
		if m.width < 40 {
			m.help = "/ filter • c cat • q quit • ↑/↓ • enter"
		} else if m.width < 60 {
			m.help = "/ filter • c cat • q quit • ↑/↓ nav • enter info"
		} else {
			m.help = "/ filter • c categories • q quit • ↑/↓ navigate • enter details"
		}

		return m, nil
	}

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
			if m.width < 40 {
				m.help = "Filter • ESC • ENTER"
			} else if m.width < 60 {
				m.help = "Filter • ESC cancel • ENTER select"
			} else {
				m.help = "Type to filter • ESC to cancel • ENTER to select"
			}
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
	// Calculate responsive dimensions
	containerWidth := m.width - 4
	if containerWidth < 30 {
		containerWidth = 30
	}
	if containerWidth > 100 {
		containerWidth = 100
	}

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

		categoryWidth := containerWidth / 2
		if categoryWidth < 25 {
			categoryWidth = 25
		}

		// Use minimal padding for very small screens
		padding := 1
		if m.width < 50 {
			padding = 0
		}
		return containerStyle.Width(categoryWidth).Padding(padding).Render(sb.String())
	}

	// Detailed view: show complete command information
	if m.detailed {
		detailWidth := containerWidth
		if detailWidth > 80 {
			detailWidth = 80
		}

		// Use minimal padding for very small screens
		padding := 1
		if m.width < 50 {
			padding = 0
		}
		detailBox := containerStyle.Width(detailWidth).Padding(padding)

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

		// Examples section with word wrapping for small screens
		var examplesText string
		maxExampleWidth := detailWidth - 8
		if padding == 0 {
			maxExampleWidth = detailWidth - 2
		}

		for _, example := range m.selectedItem.examples {
			if len(example) > maxExampleWidth && maxExampleWidth > 15 {
				// Wrap long examples on small screens
				wrapped := ""
				words := strings.Fields(example)
				currentLine := ""
				indent := "      "
				if m.width < 50 {
					indent = "  "
				}

				for _, word := range words {
					if len(currentLine)+len(word)+1 <= maxExampleWidth {
						if currentLine == "" {
							currentLine = word
						} else {
							currentLine += " " + word
						}
					} else {
						if wrapped != "" {
							wrapped += "\n" + indent
						}
						wrapped += currentLine
						currentLine = word
					}
				}
				if wrapped != "" {
					wrapped += "\n" + indent
				}
				wrapped += currentLine
				examplesText += exampleStyle.Render(wrapped) + "\n"
			} else {
				examplesText += exampleStyle.Render(example) + "\n"
			}
		}

		// Compact layout for very small screens
		if m.width < 50 {
			return detailBox.Render(
				lipgloss.JoinVertical(lipgloss.Left,
					title,
					category,
					"",
					description,
					"",
					"Usage:",
					usage,
					"",
					"Examples:",
					examplesText,
					statusMessageStyle.Render("ESC = back"),
				),
			)
		} else {
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
	}

	// List view: show command list and help message
	listContainerWidth := containerWidth
	if listContainerWidth < 30 {
		listContainerWidth = 30
	}

	// Use minimal padding for very small screens
	padding := 1
	if m.width < 50 {
		padding = 0
	}

	return containerStyle.Width(listContainerWidth).Padding(padding).Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.list.View(),
			statusMessageStyle.Render(m.help),
		),
	)
}

// ShowHelpStatic displays help information in a simple, static format
func ShowHelpStatic() {
	// Header
	fmt.Println(headingStyle.Render("R2D2 Language - Help"))
	fmt.Println()

	// Group commands by category
	categoryMap := make(map[string][]Command)
	for _, cmd := range commands {
		categoryMap[cmd.category] = append(categoryMap[cmd.category], cmd)
	}

	// Display commands by category
	categories := []string{CategoryBasic, CategoryBuild, CategoryUtil}

	for _, category := range categories {
		if cmds, exists := categoryMap[category]; exists && len(cmds) > 0 {
			// Category header
			fmt.Println(categoryStyle.Render(strings.ToUpper(category) + " COMMANDS:"))
			fmt.Println()

			// Commands in this category
			for _, cmd := range cmds {
				// Command name and description
				fmt.Printf("  %s - %s\n",
					highlightedCmdStyle.Render(cmd.name),
					cmd.description)

				// Usage
				fmt.Printf("    Usage: %s\n",
					exampleStyle.Render(cmd.usage))

				// Examples
				if len(cmd.examples) > 0 {
					fmt.Printf("    Examples:\n")
					for _, example := range cmd.examples {
						fmt.Printf("      %s\n", exampleStyle.Render(example))
					}
				}
				fmt.Println()
			}
		}
	}

	// Footer
	fmt.Println(statusMessageStyle.Render("For interactive help, use: r2d2 help (without 'static')"))
}

func ShowHelp() {
	// Get terminal dimensions
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Fallback dimensions if terminal size can't be detected
		width, height = 80, 24
	}

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

	// Calculate responsive dimensions
	listWidth := width - 4
	if listWidth < 30 {
		listWidth = 30
	}
	if listWidth > 80 {
		listWidth = 80
	}

	listHeight := height - 8
	if listHeight < 10 {
		listHeight = 10
	}
	if listHeight > 25 {
		listHeight = 25
	}

	// Setup the list model
	delegate := NewCustomDelegate(width)
	listModel := list.New(items, delegate, listWidth, listHeight)
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

	// Set initial help message based on width
	helpMsg := "/ filter • c categories • q quit • ↑/↓ navigate • enter details"
	if width < 40 {
		helpMsg = "/ filter • c cat • q quit • ↑/↓ • enter"
	} else if width < 60 {
		helpMsg = "/ filter • c cat • q quit • ↑/↓ nav • enter info"
	}

	// Initialize and run the program
	p := tea.NewProgram(
		HelpModel{
			list:           listModel,
			help:           helpMsg,
			detailed:       false,
			categories:     categories,
			filterCategory: "",
			width:          width,
			height:         height,
		},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		r2d2Styles.ErrorMessage(err.Error())
		os.Exit(1)
	}
}
