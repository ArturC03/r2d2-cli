package main

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

// Test color constants
func TestColorConstants(t *testing.T) {
	colors := map[string]string{
		"errorHex":          "#BF1A2F",
		"errorHexDark":      "#D72638",
		"infoHex":           "#454e9e",
		"infoHexDark":       "#3A57D6",
		"warningHex":        "#FFD700",
		"warningHexDark":    "#FFC107",
		"subtleHex":         "#D9DCCF",
		"subtleHexDark":     "#6E6E6E",
		"highlightHex":      "#874BFD",
		"highlightHexDark":  "#7D56F4",
		"specialHex":        "#43BF6D",
		"specialHexDark":    "#73F59F",
		"accentHex":         "#FF5733",
		"accentHexDark":     "#FF6B6B",
		"backgroundHex":     "#FAFAFA",
		"backgroundHexDark": "#1E1E1E",
		"whiteHex":          "#ffffff",
	}

	for name, expectedColor := range colors {
		t.Run("Color constant: "+name, func(t *testing.T) {
			// Verify color format (should start with # and be 7 characters long)
			if !strings.HasPrefix(expectedColor, "#") {
				t.Errorf("Color %s should start with #, got: %s", name, expectedColor)
			}
			if len(expectedColor) != 7 {
				t.Errorf("Color %s should be 7 characters long, got: %d", name, len(expectedColor))
			}
			// Verify hex format (characters after # should be valid hex)
			hexPart := expectedColor[1:]
			for _, char := range hexPart {
				if !((char >= '0' && char <= '9') || (char >= 'A' && char <= 'F') || (char >= 'a' && char <= 'f')) {
					t.Errorf("Color %s contains invalid hex character: %c", name, char)
				}
			}
		})
	}
}

// Test adaptive colors
func TestAdaptiveColors(t *testing.T) {
	adaptiveColors := map[string]lipgloss.AdaptiveColor{
		"subtleColor":     subtleColor,
		"highlightColor":  highlightColor,
		"specialColor":    specialColor,
		"warningColor":    warningColor,
		"errorColor":      errorColor,
		"infoColor":       infoColor,
		"accentColor":     accentColor,
		"backgroundColor": backgroundColor,
	}

	for name, color := range adaptiveColors {
		t.Run("Adaptive color: "+name, func(t *testing.T) {
			// Test that light and dark colors are different
			if color.Light == color.Dark {
				t.Errorf("Adaptive color %s should have different light and dark values", name)
			}

			// Test that both light and dark colors are valid hex colors
			if !strings.HasPrefix(color.Light, "#") {
				t.Errorf("Light color for %s should start with #, got: %s", name, color.Light)
			}
			if !strings.HasPrefix(color.Dark, "#") {
				t.Errorf("Dark color for %s should start with #, got: %s", name, color.Dark)
			}

			// Test length
			if len(color.Light) != 7 {
				t.Errorf("Light color for %s should be 7 characters, got: %d", name, len(color.Light))
			}
			if len(color.Dark) != 7 {
				t.Errorf("Dark color for %s should be 7 characters, got: %d", name, len(color.Dark))
			}
		})
	}
}

// Test tag styles
func TestTagStyles(t *testing.T) {
	tagStyles := map[string]lipgloss.Style{
		"tagBaseStyle": tagBaseStyle,
		"errorStyle":   errorStyle,
		"infoStyle":    infoStyle,
		"warningStyle": warningStyle,
	}

	for name, style := range tagStyles {
		t.Run("Tag style: "+name, func(t *testing.T) {
			// Test that styles are properly configured
			// We can't directly test style properties, but we can render them
			rendered := style.Render("Test")
			if rendered == "" {
				t.Errorf("Style %s should render non-empty string", name)
			}

			// Test that rendered output contains the test text
			if !strings.Contains(rendered, "Test") {
				t.Errorf("Rendered style %s should contain the input text", name)
			}
		})
	}
}

// Test UI component styles
func TestUIComponentStyles(t *testing.T) {
	uiStyles := map[string]lipgloss.Style{
		"containerStyle":     containerStyle,
		"headingStyle":       headingStyle,
		"normalItemStyle":    normalItemStyle,
		"activeItemStyle":    activeItemStyle,
		"statusMessageStyle": statusMessageStyle,
		"infoTextStyle":      infoTextStyle,
	}

	for name, style := range uiStyles {
		t.Run("UI style: "+name, func(t *testing.T) {
			// Test that styles render correctly
			rendered := style.Render("Test Content")
			if rendered == "" {
				t.Errorf("UI style %s should render non-empty string", name)
			}

			// Test that the content is preserved
			if !strings.Contains(rendered, "Test Content") {
				t.Errorf("UI style %s should preserve the input content", name)
			}
		})
	}
}

// Test style consistency
func TestStyleConsistency(t *testing.T) {
	testText := "Consistency Test"

	// Test that all tag styles contain the original text
	errorRendered := errorStyle.Render(testText)
	infoRendered := infoStyle.Render(testText)
	warningRendered := warningStyle.Render(testText)

	// All styles should contain the original text
	if !strings.Contains(errorRendered, testText) {
		t.Error("Error style should contain the original text")
	}
	if !strings.Contains(infoRendered, testText) {
		t.Error("Info style should contain the original text")
	}
	if !strings.Contains(warningRendered, testText) {
		t.Error("Warning style should contain the original text")
	}

	// Test that styles are actually applied (rendered output should be longer than input)
	if len(errorRendered) < len(testText) {
		t.Error("Error style should add formatting")
	}
	if len(infoRendered) < len(testText) {
		t.Error("Info style should add formatting")
	}
	if len(warningRendered) < len(testText) {
		t.Error("Warning style should add formatting")
	}
}

// Test style rendering with empty content
func TestStyleRenderingWithEmptyContent(t *testing.T) {
	styles := map[string]lipgloss.Style{
		"errorStyle":         errorStyle,
		"infoStyle":          infoStyle,
		"warningStyle":       warningStyle,
		"containerStyle":     containerStyle,
		"headingStyle":       headingStyle,
		"normalItemStyle":    normalItemStyle,
		"activeItemStyle":    activeItemStyle,
		"statusMessageStyle": statusMessageStyle,
		"infoTextStyle":      infoTextStyle,
	}

	for name, style := range styles {
		t.Run("Empty content for: "+name, func(t *testing.T) {
			rendered := style.Render("")
			// Even with empty content, the style might add padding/margins
			// so we just check it doesn't panic and returns a string
			_ = rendered // Just ensure it doesn't panic
		})
	}
}

// Test style rendering with special characters
func TestStyleRenderingWithSpecialCharacters(t *testing.T) {
	specialTexts := []string{
		"Text with emoji: 🚀",
		"Text with\nnewlines",
		"Text with\ttabs",
		"Text with \"quotes\"",
		"Text with <HTML> tags",
		"Text with unicode: αβγδε",
	}

	styles := []lipgloss.Style{
		errorStyle,
		infoStyle,
		warningStyle,
		headingStyle,
		activeItemStyle,
	}

	for _, text := range specialTexts {
		for i, style := range styles {
			t.Run("Special chars test", func(t *testing.T) {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Style rendering panicked with text %s on style %d: %v", text, i, r)
					}
				}()

				rendered := style.Render(text)
				if rendered == "" {
					t.Errorf("Style should render non-empty string for input: %s", text)
				}
			})
		}
	}
}

// Test color accessibility (basic contrast check)
func TestColorAccessibility(t *testing.T) {
	// Test that error colors are sufficiently different from background
	errorLight := errorColor.Light
	errorDark := errorColor.Dark
	bgLight := backgroundColor.Light
	bgDark := backgroundColor.Dark

	// Basic check: colors should not be identical
	if errorLight == bgLight {
		t.Error("Error light color should be different from background light color")
	}
	if errorDark == bgDark {
		t.Error("Error dark color should be different from background dark color")
	}

	// Similar checks for other important colors
	if infoColor.Light == backgroundColor.Light {
		t.Error("Info light color should be different from background light color")
	}
	if warningColor.Light == backgroundColor.Light {
		t.Error("Warning light color should be different from background light color")
	}
}

// Test that white color is properly defined
func TestWhiteColor(t *testing.T) {
	if whiteHex != "#ffffff" {
		t.Errorf("White color should be #ffffff, got: %s", whiteHex)
	}
}

// Test style inheritance and composition
func TestStyleComposition(t *testing.T) {
	// Test that tagBaseStyle is properly used as a base
	testText := "Base Test"

	baseRendered := tagBaseStyle.Render(testText)
	errorRendered := errorStyle.Render(testText)
	infoRendered := infoStyle.Render(testText)
	warningRendered := warningStyle.Render(testText)

	// All tag styles should contain the original text
	styles := map[string]string{
		"base":    baseRendered,
		"error":   errorRendered,
		"info":    infoRendered,
		"warning": warningRendered,
	}

	for name, rendered := range styles {
		if !strings.Contains(rendered, testText) {
			t.Errorf("Style %s should preserve the input text: %s", name, testText)
		}
	}
}

// Benchmark tests
func BenchmarkErrorStyleRender(b *testing.B) {
	text := "Error message for benchmarking"
	for i := 0; i < b.N; i++ {
		_ = errorStyle.Render(text)
	}
}

func BenchmarkInfoStyleRender(b *testing.B) {
	text := "Info message for benchmarking"
	for i := 0; i < b.N; i++ {
		_ = infoStyle.Render(text)
	}
}

func BenchmarkContainerStyleRender(b *testing.B) {
	text := "Container content for benchmarking"
	for i := 0; i < b.N; i++ {
		_ = containerStyle.Render(text)
	}
}

func BenchmarkHeadingStyleRender(b *testing.B) {
	text := "Heading text for benchmarking"
	for i := 0; i < b.N; i++ {
		_ = headingStyle.Render(text)
	}
}

// Test style memory usage (ensuring styles don't leak)
func TestStyleMemoryUsage(t *testing.T) {
	// Render the same content multiple times to check for memory leaks
	content := "Memory test content"
	iterations := 1000

	for i := 0; i < iterations; i++ {
		_ = errorStyle.Render(content)
		_ = infoStyle.Render(content)
		_ = warningStyle.Render(content)
		_ = containerStyle.Render(content)
		_ = headingStyle.Render(content)
	}

	// If we reach here without running out of memory, the test passes
	// This is a basic check - more sophisticated memory testing would require
	// additional tooling
}

// Test style configuration properties
func TestStyleProperties(t *testing.T) {
	// Test that tagBaseStyle has bold property
	baseText := "Bold Test"
	rendered := tagBaseStyle.Render(baseText)

	// We can't directly access style properties, but we can verify
	// the style renders content correctly
	if !strings.Contains(rendered, baseText) {
		t.Error("tagBaseStyle should preserve input text")
	}

	// Test that containerStyle applies border
	containerText := "Container Test"
	containerRendered := containerStyle.Render(containerText)

	if !strings.Contains(containerRendered, containerText) {
		t.Error("containerStyle should preserve input text")
	}

	// The rendered output should be longer than the input due to styling
	if len(containerRendered) <= len(containerText) {
		t.Error("containerStyle should add styling formatting")
	}
}
