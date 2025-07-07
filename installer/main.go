package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	version = "1.0.0"
)

type installState int

const (
	stateWelcome installState = iota
	stateDetectingOS
	stateInstallingDeps
	stateCloning
	stateBuilding
	stateInstalling
	stateCompleted
	stateError
)

type model struct {
	state        installState
	spinner      spinner.Model
	progress     progress.Model
	message      string
	error        error
	osInfo       OSInfo
	installPath  string
	currentStep  int
	totalSteps   int
	stepMessages []string
	width        int
	height       int
}

type OSInfo struct {
	OS         string
	PkgManager string
	BinDir     string
}

func initialModel() model {
	// Initialize styles first
	initStyles()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	p := progress.New(progress.WithDefaultGradient())
	p.Width = 30

	return model{
		state:       stateWelcome,
		spinner:     s,
		progress:    p,
		totalSteps:  6,
		currentStep: 0,
		stepMessages: []string{
			"Detecting OS...",
			"Installing deps...",
			"Cloning repo...",
			"Building CLI...",
			"Installing...",
			"Finalizing...",
		},
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		tea.EnterAltScreen,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.state == stateWelcome {
				m.state = stateDetectingOS
				m.currentStep = 1
				return m, tea.Batch(
					m.spinner.Tick,
					detectOS(),
				)
			}
			if m.state == stateCompleted || m.state == stateError {
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.progress.Width = msg.Width - 8
		if m.progress.Width > 40 {
			m.progress.Width = 40
		}
		if m.progress.Width < 15 {
			m.progress.Width = 15
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case osDetectedMsg:
		m.osInfo = OSInfo(msg)
		m.state = stateInstallingDeps
		m.currentStep = 2
		return m, installDependencies(m.osInfo)

	case depsInstalledMsg:
		m.state = stateCloning
		m.currentStep = 3
		return m, cloneRepository()

	case repositoryClonedMsg:
		m.state = stateBuilding
		m.currentStep = 4
		return m, buildCLI(string(msg))

	case cliBuiltMsg:
		m.state = stateInstalling
		m.currentStep = 5
		return m, installBinary(string(msg), m.osInfo)

	case binaryInstalledMsg:
		m.installPath = string(msg)
		m.state = stateCompleted
		m.currentStep = 6
		return m, nil

	case errorMsg:
		m.error = error(msg)
		m.state = stateError
		return m, nil

	case progressMsg:
		return m, m.progress.SetPercent(float64(msg))
	}

	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	// Header with R2D2 ASCII art (responsive)
	var header string
	if m.width > 70 {
		// Full ASCII art
		header = lipgloss.NewStyle().
			Foreground(r2d2Purple).Render("██████╗ ") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("██████╗ ") +
			lipgloss.NewStyle().
				Foreground(r2d2Green).Render("██████╗ ") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("██████╗") + "\n" +
			lipgloss.NewStyle().
				Foreground(r2d2Purple).Render("██╔══██╗") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("╚════██╗") +
			lipgloss.NewStyle().
				Foreground(r2d2Green).Render("██╔══██╗") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("╚════██╗") + "\n" +
			lipgloss.NewStyle().
				Foreground(r2d2Purple).Render("██████╔╝") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render(" █████╔╝") +
			lipgloss.NewStyle().
				Foreground(r2d2Green).Render("██║  ██║") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render(" █████╔╝") + "\n" +
			lipgloss.NewStyle().
				Foreground(r2d2Purple).Render("██╔══██╗") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("██╔═══╝ ") +
			lipgloss.NewStyle().
				Foreground(r2d2Green).Render("██║  ██║") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("██╔═══╝") + "\n" +
			lipgloss.NewStyle().
				Foreground(r2d2Purple).Render("██║  ██║") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("███████╗") +
			lipgloss.NewStyle().
				Foreground(r2d2Green).Render("██████╔╝") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("███████╗") + "\n" +
			lipgloss.NewStyle().
				Foreground(r2d2Purple).Render("╚═╝  ╚═╝") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("╚══════╝") +
			lipgloss.NewStyle().
				Foreground(r2d2Green).Render("╚═════╝ ") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("╚══════╝")
	} else {
		// Compact version
		header = lipgloss.NewStyle().
			Foreground(r2d2Purple).Render("R") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("2") +
			lipgloss.NewStyle().
				Foreground(r2d2Green).Render("D") +
			lipgloss.NewStyle().
				Foreground(r2d2Numbers).Render("2") +
			lipgloss.NewStyle().
				Foreground(accentColor).Render(" CLI Installer")
	}

	s.WriteString(lipgloss.NewStyle().Align(lipgloss.Center).Render(header))
	s.WriteString("\n")

	switch m.state {
	case stateWelcome:
		s.WriteString(infoStyle.Render("Welcome! This installer will:"))
		s.WriteString("\n")
		s.WriteString(listItemStyle.Render("• Detect your OS and package manager"))
		s.WriteString("\n")
		s.WriteString(listItemStyle.Render("• Install dependencies (Git, Go, Deno)"))
		s.WriteString("\n")
		s.WriteString(listItemStyle.Render("• Build and install R2D2 CLI"))
		s.WriteString("\n")
		s.WriteString(promptStyle.Render("Press Enter to begin or 'q' to quit"))

	case stateError:
		s.WriteString(errorStyle.Render("❌ Installation failed!"))
		s.WriteString("\n")
		s.WriteString(errorTextStyle.Render(fmt.Sprintf("Error: %v", m.error)))
		s.WriteString("\n")
		s.WriteString(promptStyle.Render("Press Enter to exit"))

	case stateCompleted:
		s.WriteString(successStyle.Render("🎉 Installation complete!"))
		s.WriteString("\n")
		s.WriteString(infoStyle.Render(fmt.Sprintf("Installed to: %s", m.installPath)))
		s.WriteString("\n")
		s.WriteString(subtleStyle.Render("Try: "))
		s.WriteString(codeStyle.Render("r2d2 --help"))
		s.WriteString("\n")
		s.WriteString(promptStyle.Render("Press Enter to exit"))

	default:
		// Progress view for installation steps
		progressPercent := float64(m.currentStep) / float64(m.totalSteps)
		s.WriteString(m.progress.ViewAs(progressPercent))
		s.WriteString("\n")

		// Current step indicator
		if m.currentStep > 0 && m.currentStep <= len(m.stepMessages) {
			s.WriteString(fmt.Sprintf("%s %s", m.spinner.View(), m.stepMessages[m.currentStep-1]))
		}

		s.WriteString("\n")

		// OS info if detected
		if m.osInfo.OS != "" {
			s.WriteString(infoStyle.Render(fmt.Sprintf("OS: %s (%s)", m.osInfo.OS, m.osInfo.PkgManager)))
			s.WriteString("\n")
		}

		s.WriteString(subtleStyle.Render("Press 'q' to quit"))
	}

	// Make container responsive
	container := containerStyle.Copy()
	if m.width > 0 {
		maxWidth := m.width - 6
		if maxWidth > 50 {
			maxWidth = 50
		}
		if maxWidth < 30 {
			maxWidth = 30
		}
		container = container.MaxWidth(maxWidth)
	}

	return container.Render(s.String())
}

// Messages
type osDetectedMsg OSInfo
type depsInstalledMsg struct{}
type repositoryClonedMsg string
type cliBuiltMsg string
type binaryInstalledMsg string
type errorMsg error
type progressMsg float64

// Commands
func detectOS() tea.Cmd {
	return func() tea.Msg {
		osInfo := OSInfo{}

		switch runtime.GOOS {
		case "linux":
			osInfo.OS = "Linux"
			osInfo.BinDir = "/usr/local/bin"

			// Detect package manager
			if _, err := exec.LookPath("apt-get"); err == nil {
				osInfo.PkgManager = "apt-get"
			} else if _, err := exec.LookPath("yum"); err == nil {
				osInfo.PkgManager = "yum"
			} else if _, err := exec.LookPath("pacman"); err == nil {
				osInfo.PkgManager = "pacman"
			} else if _, err := exec.LookPath("zypper"); err == nil {
				osInfo.PkgManager = "zypper"
			} else {
				osInfo.PkgManager = "unknown"
			}

		case "darwin":
			osInfo.OS = "macOS"
			osInfo.PkgManager = "brew"
			osInfo.BinDir = "/usr/local/bin"

		case "windows":
			osInfo.OS = "Windows"
			osInfo.PkgManager = "choco"
			osInfo.BinDir = filepath.Join(os.Getenv("USERPROFILE"), "bin")

		default:
			return errorMsg(fmt.Errorf("unsupported operating system: %s", runtime.GOOS))
		}

		// Check if we have write permissions to the bin directory
		if _, err := os.Stat(osInfo.BinDir); os.IsNotExist(err) {
			homeDir, _ := os.UserHomeDir()
			osInfo.BinDir = filepath.Join(homeDir, "bin")
		}

		time.Sleep(time.Second) // Simulate detection time
		return osDetectedMsg(osInfo)
	}
}

func installDependencies(osInfo OSInfo) tea.Cmd {
	return func() tea.Msg {
		// Check if dependencies are already installed
		if commandExists("git") && commandExists("go") && commandExists("deno") {
			time.Sleep(500 * time.Millisecond)
			return depsInstalledMsg{}
		}

		var cmd *exec.Cmd

		switch osInfo.PkgManager {
		case "apt-get":
			cmd = exec.Command("sudo", "apt-get", "update")
			if err := cmd.Run(); err != nil {
				return errorMsg(fmt.Errorf("failed to update package list: %v", err))
			}

			cmd = exec.Command("sudo", "apt-get", "install", "-y", "git", "golang-go")
			if err := cmd.Run(); err != nil {
				return errorMsg(fmt.Errorf("failed to install dependencies: %v", err))
			}

		case "brew":
			if !commandExists("brew") {
				return errorMsg(fmt.Errorf("homebrew is not installed. Please install it first from https://brew.sh/"))
			}

			cmd = exec.Command("brew", "install", "git", "go", "deno")
			if err := cmd.Run(); err != nil {
				return errorMsg(fmt.Errorf("failed to install dependencies: %v", err))
			}

		case "yum":
			cmd = exec.Command("sudo", "yum", "install", "-y", "git", "golang")
			if err := cmd.Run(); err != nil {
				return errorMsg(fmt.Errorf("failed to install dependencies: %v", err))
			}

		case "pacman":
			cmd = exec.Command("sudo", "pacman", "-S", "--noconfirm", "git", "go")
			if err := cmd.Run(); err != nil {
				return errorMsg(fmt.Errorf("failed to install dependencies: %v", err))
			}
		}

		// Install Deno if not already installed and not on macOS with brew
		if !commandExists("deno") && !(osInfo.OS == "macOS" && osInfo.PkgManager == "brew") {
			cmd = exec.Command("curl", "-fsSL", "https://deno.land/install.sh")
			output, err := cmd.Output()
			if err != nil {
				return errorMsg(fmt.Errorf("failed to download Deno installer: %v", err))
			}

			cmd = exec.Command("sh", "-c", string(output))
			if err := cmd.Run(); err != nil {
				return errorMsg(fmt.Errorf("failed to install Deno: %v", err))
			}
		}

		return depsInstalledMsg{}
	}
}

func cloneRepository() tea.Cmd {
	return func() tea.Msg {
		tmpDir := os.TempDir()
		repoDir := filepath.Join(tmpDir, "r2d2-cli-install")

		// Remove existing directory if it exists
		os.RemoveAll(repoDir)

		cmd := exec.Command("git", "clone", "https://github.com/ArturC03/r2d2-cli.git", repoDir)
		if err := cmd.Run(); err != nil {
			return errorMsg(fmt.Errorf("failed to clone repository: %v", err))
		}

		return repositoryClonedMsg(repoDir)
	}
}

func buildCLI(repoDir string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		cmd := exec.CommandContext(ctx, "go", "mod", "tidy")
		cmd.Dir = repoDir

		if err := cmd.Run(); err != nil {
			// Get the full command string for the error message
			executedCommand := cmd.String()
			// return errorMsg(fmt.Errorf("failed to tidy Go modules. Command executed: '%s', Error: %v", executedCommand, err))
			for {
				fmt.Printf("failed to tidy Go modules. Command executed: '%s', Error: %v\n", executedCommand, err)
				fmt.Print(cmd.Dir)
			}
		}

		binaryPath := filepath.Join(repoDir, "r2d2")
		if runtime.GOOS == "windows" {
			binaryPath += ".exe"
		}

		cmd = exec.CommandContext(ctx, "go", "build", "-o", binaryPath, ".")
		cmd.Dir = repoDir
		if err := cmd.Run(); err != nil {
			return errorMsg(fmt.Errorf("failed to build R2D2 CLI: %v", err))
		}

		return cliBuiltMsg(binaryPath)
	}
}

func installBinary(binaryPath string, osInfo OSInfo) tea.Cmd {
	return func() tea.Msg {
		// Ensure bin directory exists
		if err := os.MkdirAll(osInfo.BinDir, 0755); err != nil {
			return errorMsg(fmt.Errorf("failed to create bin directory: %v", err))
		}

		binaryName := "r2d2"
		if runtime.GOOS == "windows" {
			binaryName += ".exe"
		}

		targetPath := filepath.Join(osInfo.BinDir, binaryName)

		// Copy binary to target location
		cmd := exec.Command("cp", binaryPath, targetPath)
		if runtime.GOOS == "windows" {
			cmd = exec.Command("copy", binaryPath, targetPath)
		}

		if err := cmd.Run(); err != nil {
			return errorMsg(fmt.Errorf("failed to install binary: %v", err))
		}

		// Make binary executable on Unix systems
		if runtime.GOOS != "windows" {
			if err := os.Chmod(targetPath, 0755); err != nil {
				return errorMsg(fmt.Errorf("failed to make binary executable: %v", err))
			}
		}

		return binaryInstalledMsg(targetPath)
	}
}

// Helper functions
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func main() {
	// Check for version flag
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("R2D2 CLI Installer v%s\n", version)
		return
	}

	// Check for help flag
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Printf("R2D2 CLI Installer v%s\n", version)
		fmt.Println("A beautiful TUI installer for the R2D2 CLI tool")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  r2d2-installer [options]")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -v, --version    Show version")
		fmt.Println("  -h, --help       Show help")
		fmt.Println("")
		fmt.Println("Controls:")
		fmt.Println("  Enter            Begin installation")
		fmt.Println("  q, Ctrl+C        Quit installer")
		return
	}

	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running installer: %v\n", err)
		os.Exit(1)
	}
}
