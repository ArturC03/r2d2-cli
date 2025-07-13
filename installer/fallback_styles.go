package main

import (
	"github.com/charmbracelet/lipgloss"
)

// Fallback color palette
var (
	// R2D2 Brand colors
	r2d2Purple  = lipgloss.Color("#8B5CF6")                                 // Purple for R
	r2d2Green   = lipgloss.Color("#22C55E")                                 // Green for D
	r2d2Numbers = lipgloss.AdaptiveColor{Light: "#374151", Dark: "#D1D5DB"} // Adaptive for 2s

	// Primary colors
	primaryRed    = lipgloss.Color("#DC3545")
	primaryYellow = lipgloss.Color("#FFC107")
	primaryBlue   = lipgloss.Color("#3B82F6")

	// Text colors
	textPrimary   = lipgloss.AdaptiveColor{Light: "#374151", Dark: "#D1D5DB"} // Adaptive for 2s
	textSecondary = lipgloss.Color("#6C757D")
	textMuted     = lipgloss.Color("#ADB5BD")

	// Background colors
	bgLight = lipgloss.Color("#F8F9FA")
	bgDark  = lipgloss.Color("#1F2937")

	// Adaptive colors
	accentColor    = lipgloss.AdaptiveColor{Light: string(r2d2Purple), Dark: "#A78BFA"}
	successColor   = lipgloss.AdaptiveColor{Light: string(r2d2Green), Dark: "#4ADE80"}
	errorColor     = lipgloss.AdaptiveColor{Light: string(primaryRed), Dark: "#F87171"}
	warningColor   = lipgloss.AdaptiveColor{Light: string(primaryYellow), Dark: "#FBBF24"}
	infoColor      = lipgloss.AdaptiveColor{Light: string(primaryBlue), Dark: "#60A5FA"}
	subtleColor    = lipgloss.AdaptiveColor{Light: string(textSecondary), Dark: "#9CA3AF"}
	highlightColor = lipgloss.AdaptiveColor{Light: string(r2d2Purple), Dark: "#A78BFA"}
	borderColor    = lipgloss.AdaptiveColor{Light: "#E5E7EB", Dark: "#4B5563"}
	bgColor        = lipgloss.AdaptiveColor{Light: string(bgLight), Dark: string(bgDark)}
)

// Fallback styles - these will be used if r2d2Styles is not available
var (
	// Container styles
	fallbackContainerStyle = lipgloss.NewStyle().
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(borderColor).
				MaxWidth(50)

	// Header and title styles
	fallbackHeaderStyle = lipgloss.NewStyle().
				Foreground(accentColor).
				Bold(true).
				Align(lipgloss.Center).
				MarginBottom(0)

	fallbackTitleStyle = lipgloss.NewStyle().
				Foreground(highlightColor).
				Bold(true).
				Align(lipgloss.Center).
				MarginBottom(1)

	// Status message styles
	fallbackInfoStyle = lipgloss.NewStyle().
				Foreground(infoColor).
				Bold(true).
				Padding(0, 1)

	fallbackSuccessStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true).
				Padding(0, 1)

	fallbackErrorStyle = lipgloss.NewStyle().
				Foreground(errorColor).
				Bold(true).
				Padding(0, 1)

	fallbackWarningStyle = lipgloss.NewStyle().
				Foreground(warningColor).
				Bold(true).
				Padding(0, 1)

	// Text styles
	fallbackSubtleStyle = lipgloss.NewStyle().
				Foreground(subtleColor).
				MarginBottom(0)

	fallbackListItemStyle = lipgloss.NewStyle().
				Foreground(textPrimary).
				PaddingLeft(1)

	fallbackErrorTextStyle = lipgloss.NewStyle().
				Foreground(errorColor).
				Italic(true).
				PaddingLeft(1)

	// Interactive styles
	fallbackPromptStyle = lipgloss.NewStyle().
				Foreground(accentColor).
				Bold(true).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(accentColor).
				Padding(0, 1).
				MarginTop(0)

	fallbackCodeStyle = lipgloss.NewStyle().
				Foreground(highlightColor).
				Background(bgColor).
				Padding(0, 1).
				MarginLeft(1)

	// Spinner style
	fallbackSpinnerStyle = lipgloss.NewStyle().
				Foreground(accentColor)
)
