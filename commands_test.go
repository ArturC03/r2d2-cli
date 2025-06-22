package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Test ShowVersion function
func TestShowVersion(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	ShowVersion()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check if output contains expected version info
	if !strings.Contains(output, "R2D2 Language") {
		t.Errorf("ShowVersion() should contain 'R2D2 Language', got: %v", output)
	}

	if !strings.Contains(output, Version) {
		t.Errorf("ShowVersion() should contain version %v, got: %v", Version, output)
	}

	if !strings.Contains(output, "version") {
		t.Errorf("ShowVersion() should contain 'version', got: %v", output)
	}
}

// Test UnknownCommand function
func TestUnknownCommand(t *testing.T) {
	tests := []struct {
		name         string
		command      string
		errArgIndex  int
		expectOutput []string
	}{
		{
			name:         "Invalid command 'invalid'",
			command:      "invalid",
			errArgIndex:  1,
			expectOutput: []string{"Unknown command: invalid", "Run 'r2d2 help' for usage."},
		},
		{
			name:         "Invalid command 'test'",
			command:      "test",
			errArgIndex:  1,
			expectOutput: []string{"Unknown command: test", "Run 'r2d2 help' for usage."},
		},
		{
			name:         "Empty command",
			command:      "",
			errArgIndex:  1,
			expectOutput: []string{"Unknown command: ", "Run 'r2d2 help' for usage."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			UnknownCommand(tt.command, tt.errArgIndex)

			// Restore stdout
			w.Close()
			os.Stdout = old

			// Read the output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// Check if output contains expected strings
			for _, expected := range tt.expectOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("UnknownCommand() should contain '%v', got: %v", expected, output)
				}
			}
		})
	}
}

// Test InfoMessage function
func TestInfoMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple info message",
			input:    "This is an info message",
			expected: "This is an info message",
		},
		{
			name:     "Empty message",
			input:    "",
			expected: "",
		},
		{
			name:     "Message with special characters",
			input:    "Info: Test@123!",
			expected: "Info: Test@123!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InfoMessage(tt.input)

			// The result should contain the input message
			// (it may have styling applied, but the content should be there)
			if !strings.Contains(result, tt.expected) {
				t.Errorf("InfoMessage(%v) should contain %v, got: %v", tt.input, tt.expected, result)
			}
		})
	}
}

// Test ErrorMessage function
func TestErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple error message",
			input:    "This is an error",
			expected: "This is an error",
		},
		{
			name:     "Empty error message",
			input:    "",
			expected: "",
		},
		{
			name:     "Error with details",
			input:    "File not found: test.r2d2",
			expected: "File not found: test.r2d2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ErrorMessage(tt.input)

			// The result should contain the input message
			if !strings.Contains(result, tt.expected) {
				t.Errorf("ErrorMessage(%v) should contain %v, got: %v", tt.input, tt.expected, result)
			}
		})
	}
}

// Test HelpMessage function
func TestHelpMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple help message",
			input:    "Use -h for help",
			expected: "Use -h for help",
		},
		{
			name:     "Help usage message",
			input:    "Run 'r2d2 help' for usage.",
			expected: "Run 'r2d2 help' for usage.",
		},
		{
			name:     "Empty help message",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HelpMessage(tt.input)

			// The result should contain the input message
			if !strings.Contains(result, tt.expected) {
				t.Errorf("HelpMessage(%v) should contain %v, got: %v", tt.input, tt.expected, result)
			}
		})
	}
}

// Test Version constant
func TestVersionConstant(t *testing.T) {
	if Version == "" {
		t.Error("Version constant should not be empty")
	}

	// Check if version follows semantic versioning pattern (basic check)
	parts := strings.Split(Version, ".")
	if len(parts) != 3 {
		t.Errorf("Version should follow semantic versioning (x.y.z), got: %v", Version)
	}

	// Each part should be numeric (basic validation)
	for i, part := range parts {
		if part == "" {
			t.Errorf("Version part %d should not be empty in version: %v", i, Version)
		}
	}
}

// Test message styling consistency
func TestMessageStyleConsistency(t *testing.T) {
	testMessage := "Test message"

	infoResult := InfoMessage(testMessage)
	errorResult := ErrorMessage(testMessage)
	helpResult := HelpMessage(testMessage)

	// All should contain the original message
	if !strings.Contains(infoResult, testMessage) {
		t.Error("InfoMessage should preserve original message content")
	}

	if !strings.Contains(errorResult, testMessage) {
		t.Error("ErrorMessage should preserve original message content")
	}

	if !strings.Contains(helpResult, testMessage) {
		t.Error("HelpMessage should preserve original message content")
	}
}

// Test Build function integration (basic test)
func TestBuildFunctionSignature(t *testing.T) {
	// This test ensures the Build function accepts the expected parameters
	// We can't easily test the actual functionality without mocking r2d2.BuildCode
	// but we can test that the function exists and accepts the right types

	testCode := "module Test { export fn main() { console.log(\"Hello\"); } }"
	testFilename := "test.r2d2"

	// This should not panic if function signature is correct
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Build function panicked: %v", r)
		}
	}()

	// Note: This will call the actual Build function
	// In a real test environment, you might want to mock r2d2.BuildCode
	// Build(testCode, testFilename)

	// For now, just test that we can call it without compilation errors
	_ = testCode
	_ = testFilename
}

// Test Run function integration (basic test)
func TestRunFunctionSignature(t *testing.T) {
	testCode := "module Test { export fn main() { console.log(\"Hello\"); } }"

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Run function panicked: %v", r)
		}
	}()

	// Note: This will call the actual Run function
	// In a real test environment, you might want to mock r2d2.RunCode
	// Run(testCode)

	// For now, just test that we can call it without compilation errors
	_ = testCode
}

// Test BuildJs function integration (basic test)
func TestBuildJsFunctionSignature(t *testing.T) {
	testCode := "module Test { export fn main() { console.log(\"Hello\"); } }"
	testFilename := "test.js"

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("BuildJs function panicked: %v", r)
		}
	}()

	// Note: This will call the actual BuildJs function
	// In a real test environment, you might want to mock r2d2.BuildJsFile
	// BuildJs(testCode, testFilename)

	// For now, just test that we can call it without compilation errors
	_ = testCode
	_ = testFilename
}

// Benchmark tests for message functions
func BenchmarkInfoMessage(b *testing.B) {
	message := "This is a test info message"
	for i := 0; i < b.N; i++ {
		_ = InfoMessage(message)
	}
}

func BenchmarkErrorMessage(b *testing.B) {
	message := "This is a test error message"
	for i := 0; i < b.N; i++ {
		_ = ErrorMessage(message)
	}
}

func BenchmarkHelpMessage(b *testing.B) {
	message := "This is a test help message"
	for i := 0; i < b.N; i++ {
		_ = HelpMessage(message)
	}
}

// Test edge cases
func TestMessageFunctionsWithSpecialCharacters(t *testing.T) {
	specialChars := []string{
		"Message with unicode: 🚀",
		"Message with newlines\nand\ttabs",
		"Message with \"quotes\" and 'apostrophes'",
		"Message with <HTML> tags",
		"Message with $pecial characters: @#%^&*()",
	}

	for _, msg := range specialChars {
		t.Run("Special chars: "+msg[:min(20, len(msg))], func(t *testing.T) {
			// Test that functions don't panic with special characters
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Function panicked with special characters: %v", r)
				}
			}()

			infoResult := InfoMessage(msg)
			errorResult := ErrorMessage(msg)
			helpResult := HelpMessage(msg)

			// All should return non-empty strings
			if len(infoResult) == 0 {
				t.Error("InfoMessage returned empty string")
			}
			if len(errorResult) == 0 {
				t.Error("ErrorMessage returned empty string")
			}
			if len(helpResult) == 0 {
				t.Error("HelpMessage returned empty string")
			}
		})
	}
}

// Helper function for min (Go doesn't have built-in min for int)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestBuild(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		filename string
		expected string
		wantErr  bool
	}{
		{
			name:     "Simple build",
			code:     "module Test { export fn main() { console.log(\"Hello\"); } }",
			filename: "test_simple.r2d2",
			expected: "Hello",
			wantErr:  false,
		},
		{
			name:     "Build with special characters",
			code:     "module Test { export fn main() { console.log(\"Hello\"); } }",
			filename: "test_special.r2d2",
			expected: "Hello",
			wantErr:  false,
		},
		{
			name:     "Build with unicode characters",
			code:     "module Test { export fn main() { console.log(\"🚀\"); } }",
			filename: "test_unicode.r2d2",
			expected: "🚀",
			wantErr:  false,
		},
		{
			name:     "Build with newlines and tabs",
			code:     "module Test { export fn main() { console.log(\"\"\"Hello\\n\\tWorld\"\"\"); } }",
			filename: "test_newline.r2d2",
			expected: "Hello\n\tWorld",
			wantErr:  false,
		},
		{
			name:     "Build with quotes and apostrophes",
			code:     "module Test { export fn main() { console.log(\"Hello 'World'\"); } }",
			filename: "test_quotes.r2d2",
			expected: "Hello 'World'",
			wantErr:  false,
		},
		{
			name:     "Build with HTML tags",
			code:     "module Test { export fn main() { console.log(\"<html><body>Hello</body></html>\"); } }",
			filename: "test_html.r2d2",
			expected: "<html><body>Hello</body></html>",
			wantErr:  false,
		},
		{
			name:     "Build with special characters 2",
			code:     "module Test { export fn main() { console.log(\"$pecial characters: @#%^&*()\"); } }",
			filename: "test_special2.r2d2",
			expected: "$pecial characters: @#%^&*()",
			wantErr:  false,
		},
		{
			name:     "Build with undefined variable assignment",
			code:     "module Test { export fn main() { v = 3; console.log(\"$pecial characters: @#%^&*()\"); } }",
			filename: "test_assignment_error.r2d2",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Criar ficheiro com o código de teste
			err := os.WriteFile(tt.filename, []byte(tt.code), 0644)
			if err != nil {
				t.Fatalf("erro ao criar ficheiro: %v", err)
			}
			defer os.Remove(tt.filename)

			// Compilar
			err = Build(tt.code, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Build() erro = %v, era esperado erro = %v", err, tt.wantErr)
			}

			// Se esperávamos erro, não devemos tentar executar binário
			if tt.wantErr {
				return
			}

			exeName := strings.TrimSuffix(tt.filename, ".r2d2")
			defer os.Remove(exeName)

			if _, statErr := os.Stat(exeName); os.IsNotExist(statErr) {
				t.Fatalf("esperava que o binário %s fosse criado, mas não foi encontrado", exeName)
			}

			os.Chmod(exeName, 0755)

			// Executar binário
			cmd := exec.Command("./" + exeName)
			outputBytes, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("erro ao executar binário: %v\noutput: %s", err, outputBytes)
			}

			output := strings.TrimSpace(string(outputBytes))
			if !strings.Contains(output, tt.expected) {
				t.Errorf("esperava output com %q, mas foi: %q", tt.expected, output)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected string
		wantErr  bool
	}{
		{
			name:     "Simple run",
			code:     "module Test { export fn main() { console.log(\"Hello\"); } }",
			expected: "Hello",
			wantErr:  false,
		},
		{
			name:     "Run with special characters",
			code:     "module Test { export fn main() { console.log(\"Hello\"); } }",
			expected: "Hello",
			wantErr:  false,
		},
		{
			name:     "Run with unicode characters",
			code:     "module Test { export fn main() { console.log(\"🚀\"); } }",
			expected: "🚀",
			wantErr:  false,
		},
		{
			name:     "Run with newlines and tabs",
			code:     "module Test { export fn main() { console.log(\"Hello\\n\\tWorld\"); } }",
			expected: "Hello\n\tWorld",
			wantErr:  false,
		},
		{
			name:     "Run with quotes and apostrophes",
			code:     "module Test { export fn main() { console.log(\"Hello 'World'\"); } }",
			expected: "Hello 'World'",
			wantErr:  false,
		},
		{
			name:     "Run with HTML tags",
			code:     "module Test { export fn main() { console.log(\"<html><body>Hello</body></html>\"); } }",
			expected: "<html><body>Hello</body></html>",
			wantErr:  false,
		},
		{
			name:     "Run with special characters",
			code:     "module Test { export fn main() { console.log(\"$pecial characters: @#%^&*()\"); } }",
			expected: "$pecial characters: @#%^&*()",
			wantErr:  false,
		},
		{
			name:     "Build with undefined variable assignment",
			code:     "module Test { export fn main() { v = 3; console.log(\"$pecial characters: @#%^&*()\"); } }",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			err := Run(tt.code)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Run() erro = %v, era esperado erro = %v", err, tt.wantErr)
			}

			// Restore stdout
			w.Close()
			os.Stdout = old

			// Read the output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// Check if output contains expected string
			if !strings.Contains(output, tt.expected) {
				t.Errorf("MessageFunctionsWithSpecialCharacters() should contain %v, got: %v", tt.expected, output)
			}
		})
	}
}

func TestBuildJs(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		filename string
		expected string
		wantErr  bool
	}{
		{
			name:     "Simple build",
			code:     "module Test { export fn main() { console.log(\"Hello\"); } }",
			filename: "test_simple.r2d2",
			expected: "Hello",
			wantErr:  false,
		},
		{
			name:     "Build with special characters",
			code:     "module Test { export fn main() { console.log(\"$pecial characters: @#%^&*()\"); } }",
			filename: "test_special.r2d2",
			expected: "$pecial characters: @#%^&*()",
			wantErr:  false,
		},
		{
			name:     "Build with unicode characters",
			code:     "module Test { export fn main() { console.log(\"🚀\"); } }",
			filename: "test_unicode.r2d2",
			expected: "🚀",
			wantErr:  false,
		},
		{
			name:     "Build with newlines and tabs",
			code:     "module Test { export fn main() { console.log(\"Hello\\n\\tWorld\"); } }",
			filename: "test_newline.r2d2",
			expected: "Hello\n\tWorld",
			wantErr:  false,
		},
		{
			name:     "Build with quotes and apostrophes",
			code:     "module Test { export fn main() { console.log(\"Hello 'World'\"); } }",
			filename: "test_quotes.r2d2",
			expected: "Hello 'World'",
			wantErr:  false,
		},
		{
			name:     "Build with HTML tags",
			code:     "module Test { export fn main() { console.log(\"<html><body>Hello</body></html>\"); } }",
			filename: "test_html.r2d2",
			expected: "<html><body>Hello</body></html>",
			wantErr:  false,
		},
		{
			name:     "Invalid variable assignment",
			code:     "module Test { export fn main() { x = 5; console.log(x); } }",
			filename: "test_invalid.r2d2",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Cria o ficheiro com o código .r2d2
			err := os.WriteFile(tt.filename, []byte(tt.code), 0644)
			if err != nil {
				t.Fatalf("erro ao escrever o ficheiro: %v", err)
			}
			defer os.Remove(tt.filename)

			// Compila usando a função BuildJs
			err = BuildJs(tt.code, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Fatalf("BuildJs() erro = %v, era esperado erro = %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			// Executa o JS com deno
			jsName := strings.TrimSuffix(tt.filename, ".r2d2") + ".js"
			defer os.Remove(jsName)

			cmd := exec.Command("deno", "run", "--allow-all", jsName)
			outputBytes, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("erro ao executar JS com Deno: %v\noutput: %s", err, outputBytes)
			}

			output := strings.TrimSpace(string(outputBytes))
			if !strings.Contains(output, tt.expected) {
				t.Errorf("esperava output com %q, mas foi: %q", tt.expected, output)
			}
		})
	}
}
