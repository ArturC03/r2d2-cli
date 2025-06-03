package main

import "github.com/charmbracelet/lipgloss"

// Define color constants
const (
	errorHex          = "#BF1A2F"
	errorHexDark      = "#D72638"
	infoHex           = "#454e9e"
	infoHexDark       = "#3A57D6"
	warningHex        = "#FFD700"
	warningHexDark    = "#FFC107"
	subtleHex         = "#D9DCCF"
	subtleHexDark     = "#6E6E6E"
	highlightHex      = "#874BFD"
	highlightHexDark  = "#7D56F4"
	specialHex        = "#43BF6D"
	specialHexDark    = "#73F59F"
	accentHex         = "#FF5733"
	accentHexDark     = "#FF6B6B"
	backgroundHex     = "#FAFAFA"
	backgroundHexDark = "#1E1E1E"
	whiteHex          = "#ffffff"
)

// Adaptive colors for theme support
var (
	subtleColor     = lipgloss.AdaptiveColor{Light: subtleHex, Dark: subtleHexDark}
	highlightColor  = lipgloss.AdaptiveColor{Light: highlightHex, Dark: highlightHexDark}
	specialColor    = lipgloss.AdaptiveColor{Light: specialHex, Dark: specialHexDark}
	warningColor    = lipgloss.AdaptiveColor{Light: warningHex, Dark: warningHexDark}       // Yellow for warnings
	errorColor      = lipgloss.AdaptiveColor{Light: errorHex, Dark: errorHexDark}           // Red for errors
	infoColor       = lipgloss.AdaptiveColor{Light: infoHex, Dark: infoHexDark}             // Blue for information
	accentColor     = lipgloss.AdaptiveColor{Light: accentHex, Dark: accentHexDark}         // Orange for highlighting
	backgroundColor = lipgloss.AdaptiveColor{Light: backgroundHex, Dark: backgroundHexDark} // Adaptive background
)

// Tag styles with consistent formatting
var (
	tagBaseStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(whiteHex)).
			Padding(0, 1)

	errorStyle = tagBaseStyle.Copy().
			Background(lipgloss.Color(errorHex))

	infoStyle = tagBaseStyle.Copy().
			Background(lipgloss.Color(infoHex))

	warningStyle = tagBaseStyle.Copy().
			Background(lipgloss.Color(warningHex)).
			Foreground(lipgloss.Color("#333333")) // Darker text for better contrast on yellow
)

// UI component styles
var (
	// Container style
	containerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(subtleColor).
			Padding(1, 2)

	// Text styles
	headingStyle = lipgloss.NewStyle().
			Foreground(highlightColor).
			Bold(true).
			Padding(0, 1).
			MarginLeft(2)

	normalItemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color("241"))

	activeItemStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(specialColor).
			Bold(true)

	statusMessageStyle = lipgloss.NewStyle().
				Italic(true).
				Foreground(subtleColor).
				PaddingTop(1).
				PaddingLeft(2)

	infoTextStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(infoColor).
			Bold(true)
)
