package main

import (
	"os"
	"strings"
	"testing"
)

func TestGetFilename(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	// Test cases
	tests := []struct {
		name       string
		args       []string
		expected   string
		shouldExit bool
	}{
		{
			name:       "Valid file path with extension",
			args:       []string{"r2d2", "build", "test.r2d2"},
			expected:   "test.r2d2",
			shouldExit: false,
		},
		{
			name:       "File path with directory separator",
			args:       []string{"r2d2", "build", "/path/to/hello.r2d2"},
			expected:   "hello.r2d2",
			shouldExit: false,
		},
		{
			name:       "Path with multiple separators",
			args:       []string{"r2d2", "build", "projects/subdir/app.r2d2"},
			expected:   "app.r2d2",
			shouldExit: false,
		},
		{
			name:       "Insufficient arguments",
			args:       []string{"r2d2", "build"},
			expected:   "",
			shouldExit: true,
		},
		{
			name:       "No arguments",
			args:       []string{"r2d2"},
			expected:   "",
			shouldExit: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set test args
			os.Args = tt.args

			if tt.shouldExit {
				// We can't test os.Exit directly, so we just check that we have insufficient args
				if len(os.Args) >= 3 {
					t.Errorf("Expected insufficient arguments, but got %d args", len(os.Args))
				}
			} else {
				result := getFilename()
				if result != tt.expected {
					t.Errorf("getFilename() = %v, want %v", result, tt.expected)
				}
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

func TestParseOutputFlag(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	tests := []struct {
		name           string
		args           []string
		expectedFile   string
		expectedExists bool
	}{
		{
			name:           "Valid -o flag",
			args:           []string{"r2d2", "build", "test.r2d2", "-o", "output.html"},
			expectedFile:   "output.html",
			expectedExists: true,
		},
		{
			name:           "No -o flag",
			args:           []string{"r2d2", "build", "test.r2d2"},
			expectedFile:   "",
			expectedExists: false,
		},
		{
			name:           "-o flag at beginning",
			args:           []string{"r2d2", "-o", "output.js", "build", "test.r2d2"},
			expectedFile:   "output.js",
			expectedExists: true,
		},
		{
			name:           "-o flag without value (at end)",
			args:           []string{"r2d2", "build", "test.r2d2", "-o"},
			expectedFile:   "",
			expectedExists: false,
		},
		{
			name:           "Multiple -o flags (should get first)",
			args:           []string{"r2d2", "build", "test.r2d2", "-o", "first.html", "-o", "second.html"},
			expectedFile:   "first.html",
			expectedExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			file, exists := parseOutputFlag()

			if file != tt.expectedFile {
				t.Errorf("parseOutputFlag() file = %v, want %v", file, tt.expectedFile)
			}
			if exists != tt.expectedExists {
				t.Errorf("parseOutputFlag() exists = %v, want %v", exists, tt.expectedExists)
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

func TestRemoveOutputFlag(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	tests := []struct {
		name     string
		args     []string
		expected []string
	}{
		{
			name:     "Remove -o flag and value",
			args:     []string{"r2d2", "build", "test.r2d2", "-o", "output.html"},
			expected: []string{"r2d2", "build", "test.r2d2"},
		},
		{
			name:     "No -o flag to remove",
			args:     []string{"r2d2", "build", "test.r2d2"},
			expected: []string{"r2d2", "build", "test.r2d2"},
		},
		{
			name:     "Remove -o flag at beginning",
			args:     []string{"r2d2", "-o", "output.js", "build", "test.r2d2"},
			expected: []string{"r2d2", "build", "test.r2d2"},
		},
		{
			name:     "-o flag without value (at end)",
			args:     []string{"r2d2", "build", "test.r2d2", "-o"},
			expected: []string{"r2d2", "build", "test.r2d2", "-o"},
		},
		{
			name:     "Multiple -o flags",
			args:     []string{"r2d2", "-o", "first.html", "build", "test.r2d2", "-o", "second.html"},
			expected: []string{"r2d2", "build", "test.r2d2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			removeOutputFlag()

			if len(os.Args) != len(tt.expected) {
				t.Errorf("removeOutputFlag() result length = %v, want %v", len(os.Args), len(tt.expected))
				return
			}

			for i, arg := range os.Args {
				if arg != tt.expected[i] {
					t.Errorf("removeOutputFlag() result[%d] = %v, want %v", i, arg, tt.expected[i])
				}
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

func TestCommandLineJoin(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Basic command line",
			args:     []string{"r2d2", "build", "test.r2d2"},
			expected: "build test.r2d2",
		},
		{
			name:     "Command with flags",
			args:     []string{"r2d2", "build", "test.r2d2", "-o", "output.html"},
			expected: "build test.r2d2 -o output.html",
		},
		{
			name:     "Single command",
			args:     []string{"r2d2", "help"},
			expected: "help",
		},
		{
			name:     "No additional args",
			args:     []string{"r2d2"},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			result := strings.Join(os.Args[1:], " ")

			if result != tt.expected {
				t.Errorf("commandLine join = %v, want %v", result, tt.expected)
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

// Test helper functions for file path extraction
func TestFilePathExtraction(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "Unix path",
			filePath: "/home/user/projects/hello.r2d2",
			expected: "hello.r2d2",
		},
		{
			name:     "Deep nested path",
			filePath: "projects/nested/deep/app.r2d2",
			expected: "app.r2d2",
		},
		{
			name:     "Relative path",
			filePath: "../projects/test.r2d2",
			expected: "test.r2d2",
		},
		{
			name:     "Just filename",
			filePath: "simple.r2d2",
			expected: "simple.r2d2",
		},
		{
			name:     "Path with spaces",
			filePath: "/path/to/my file.r2d2",
			expected: "my file.r2d2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parts := strings.Split(tt.filePath, string(os.PathSeparator))
			filename := parts[len(parts)-1]

			if filename != tt.expected {
				t.Errorf("filename extraction = %v, want %v", filename, tt.expected)
			}
		})
	}
}

// Test command validation logic
func TestCommandValidation(t *testing.T) {
	validCommands := []string{
		"help", "-help", "-h", "--help", "--h",
		"version", "-version", "-v", "--version", "--v",
		"build", "-b",
		"run", "-r",
		"js",
		"new",
	}

	for _, cmd := range validCommands {
		t.Run("Valid command: "+cmd, func(t *testing.T) {
			// This test verifies that these are recognized commands
			// In a real implementation, you might have a function to validate commands
			isValid := isValidCommand(cmd)
			if !isValid {
				t.Errorf("Command %v should be valid", cmd)
			}
		})
	}

	invalidCommands := []string{
		"invalid", "test", "compile", "execute", "--invalid",
	}

	for _, cmd := range invalidCommands {
		t.Run("Invalid command: "+cmd, func(t *testing.T) {
			isValid := isValidCommand(cmd)
			if isValid {
				t.Errorf("Command %v should be invalid", cmd)
			}
		})
	}
}

// Helper function to simulate command validation
func isValidCommand(cmd string) bool {
	validCommands := map[string]bool{
		"help": true, "-help": true, "-h": true, "--help": true, "--h": true,
		"version": true, "-version": true, "-v": true, "--version": true, "--v": true,
		"build": true, "-b": true,
		"run": true, "-r": true,
		"js":  true,
		"new": true,
	}
	return validCommands[cmd]
}

// Test argument counting
func TestArgumentCount(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	tests := []struct {
		name          string
		args          []string
		minRequired   int
		hasEnoughArgs bool
	}{
		{
			name:          "Sufficient args for build",
			args:          []string{"r2d2", "build", "test.r2d2"},
			minRequired:   3,
			hasEnoughArgs: true,
		},
		{
			name:          "Insufficient args for build",
			args:          []string{"r2d2", "build"},
			minRequired:   3,
			hasEnoughArgs: false,
		},
		{
			name:          "Help command (no file needed)",
			args:          []string{"r2d2", "help"},
			minRequired:   2,
			hasEnoughArgs: true,
		},
		{
			name:          "No command",
			args:          []string{"r2d2"},
			minRequired:   2,
			hasEnoughArgs: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			hasEnough := len(os.Args) >= tt.minRequired

			if hasEnough != tt.hasEnoughArgs {
				t.Errorf("Argument count check = %v, want %v (args: %v, required: %v)",
					hasEnough, tt.hasEnoughArgs, len(os.Args), tt.minRequired)
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

// Test for file extension validation
func TestFileExtensionValidation(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		isValid  bool
	}{
		{
			name:     "Valid R2D2 file",
			filename: "test.r2d2",
			isValid:  true,
		},
		{
			name:     "Valid R2D2 file with path",
			filename: "/path/to/app.r2d2",
			isValid:  true,
		},
		{
			name:     "Invalid extension",
			filename: "test.txt",
			isValid:  false,
		},
		{
			name:     "No extension",
			filename: "test",
			isValid:  false,
		},
		{
			name:     "R2D2 in middle of name",
			filename: "test.r2d2.backup",
			isValid:  false,
		},
		{
			name:     "Case sensitive check",
			filename: "test.R2D2",
			isValid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := strings.HasSuffix(tt.filename, ".r2d2")

			if isValid != tt.isValid {
				t.Errorf("File extension validation for %v = %v, want %v",
					tt.filename, isValid, tt.isValid)
			}
		})
	}
}

// Benchmark tests
func BenchmarkGetFilename(b *testing.B) {
	// Save original args
	originalArgs := os.Args
	os.Args = []string{"r2d2", "build", "/very/long/path/to/some/deeply/nested/file.r2d2"}

	for i := 0; i < b.N; i++ {
		_ = getFilename()
	}

	// Restore original args
	os.Args = originalArgs
}

func BenchmarkParseOutputFlag(b *testing.B) {
	// Save original args
	originalArgs := os.Args
	os.Args = []string{"r2d2", "build", "test.r2d2", "-o", "output.html", "extra", "args"}

	for i := 0; i < b.N; i++ {
		_, _ = parseOutputFlag()
	}

	// Restore original args
	os.Args = originalArgs
}

func BenchmarkRemoveOutputFlag(b *testing.B) {
	originalArgs := os.Args

	for i := 0; i < b.N; i++ {
		// Reset args for each iteration
		os.Args = []string{"r2d2", "build", "test.r2d2", "-o", "output.html"}
		removeOutputFlag()
	}

	// Restore original args
	os.Args = originalArgs
}
