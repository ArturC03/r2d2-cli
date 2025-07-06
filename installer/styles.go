package main

import (
	"github.com/charmbracelet/lipgloss"
)

// Initialize styles using fallback styles
func initStyles() {
	containerStyle = fallbackContainerStyle
	headerStyle = fallbackHeaderStyle
	titleStyle = fallbackTitleStyle
	infoStyle = fallbackInfoStyle
	successStyle = fallbackSuccessStyle
	errorStyle = fallbackErrorStyle
	warningStyle = fallbackWarningStyle
	subtleStyle = fallbackSubtleStyle
	listItemStyle = fallbackListItemStyle
	errorTextStyle = fallbackErrorTextStyle
	promptStyle = fallbackPromptStyle
	codeStyle = fallbackCodeStyle
	spinnerStyle = fallbackSpinnerStyle
}

var (
	// These will be initialized by initStyles()
	containerStyle lipgloss.Style
	headerStyle    lipgloss.Style
	titleStyle     lipgloss.Style
	infoStyle      lipgloss.Style
	successStyle   lipgloss.Style
	errorStyle     lipgloss.Style
	warningStyle   lipgloss.Style
	subtleStyle    lipgloss.Style
	listItemStyle  lipgloss.Style
	errorTextStyle lipgloss.Style
	promptStyle    lipgloss.Style
	codeStyle      lipgloss.Style
	spinnerStyle   lipgloss.Style
)
