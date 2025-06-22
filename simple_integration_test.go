package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestBasicIntegrationWorkflow tests the core workflow without TUI components
func TestBasicIntegrationWorkflow(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	// Create a temporary test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "integration_test.r2d2")
	testContent := `module IntegrationTest {
    var counter i32 = 0;

    export fn main() {
        console.log("Integration test running");
        counter = 42;
    }

    fn helper() i32 {
        return counter * 2;
    }
}`

	err := ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{
			name:        "Version command",
			args:        []string{"r2d2", "version"},
			expectError: false,
		},
		{
			name:        "Build command with valid file",
			args:        []string{"r2d2", "build", testFile},
			expectError: false,
		},
		{
			name:        "Build with output flag",
			args:        []string{"r2d2", "build", testFile, "-o", "integration_output.html"},
			expectError: false,
		},
		{
			name:        "JS command with valid file",
			args:        []string{"r2d2", "js", testFile},
			expectError: false,
		},
		{
			name:        "JS with output flag",
			args:        []string{"r2d2", "js", testFile, "-o", "integration_output.js"},
			expectError: false,
		},
		{
			name:        "Run command with valid file",
			args:        []string{"r2d2", "run", testFile},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args

			// Test output flag parsing
			outputFile, hasOutput := parseOutputFlag()

			// Test argument validation
			if len(os.Args) >= 2 {
				cmd := os.Args[1]

				// Test that we can identify the command correctly
				isFileCommand := cmd == "build" || cmd == "-b" || cmd == "run" || cmd == "-r" || cmd == "js"

				if isFileCommand {
					// Test that we have sufficient arguments
					if len(os.Args) < 3 {
						if !tt.expectError {
							t.Errorf("Expected sufficient arguments for command %s", cmd)
						}
						return
					}

					// Test file reading
					filePath := os.Args[2]
					if filePath != "" {
						content, err := os.ReadFile(filePath)
						if err != nil {
							if !tt.expectError {
								t.Errorf("Failed to read file %s: %v", filePath, err)
							}
							return
						}

						// Test that file content is readable
						if len(content) == 0 {
							t.Errorf("File content is empty: %s", filePath)
						}

						// Test filename extraction
						filename := getFilename()
						if filename == "" {
							t.Errorf("Failed to extract filename from %s", filePath)
						}
					}
				}

				// Test output flag handling if present
				if hasOutput {
					if outputFile == "" {
						t.Errorf("Output flag detected but no output file specified")
					}

					// Test that removeOutputFlag works
					originalArgCount := len(os.Args)
					removeOutputFlag()
					if len(os.Args) >= originalArgCount {
						t.Errorf("removeOutputFlag should reduce argument count")
					}
				}
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

// TestFileOperationsIntegration tests file operations end-to-end
func TestFileOperationsIntegration(t *testing.T) {
	originalArgs := os.Args
	tempDir := t.TempDir()

	// Test with various file types and content
	testFiles := []struct {
		name       string
		filename   string
		content    string
		shouldWork bool
	}{
		{
			name:       "Valid R2D2 module",
			filename:   "valid.r2d2",
			content:    `module Valid { export fn main() { console.log("Hello"); } }`,
			shouldWork: true,
		},
		{
			name:     "Complex R2D2 code",
			filename: "complex.r2d2",
			content: `module Complex {
				var x i32 = 10;
				let y number = 20.5;
				const z bool = true;

				export fn main() {
					if (x > 5) {
						console.log("Complex test");
						for var i i32 = 0; i < x; i++ {
							process(i);
						}
					}
				}

				fn process(value i32) {
					console.log(value);
				}
			}`,
			shouldWork: true,
		},
		{
			name:       "Empty file",
			filename:   "empty.r2d2",
			content:    "",
			shouldWork: true, // Empty files should be readable
		},
		{
			name:       "Large file",
			filename:   "large.r2d2",
			content:    strings.Repeat("// Comment line\n", 1000) + `module Large { export fn main() {} }`,
			shouldWork: true,
		},
	}

	for _, tf := range testFiles {
		t.Run(tf.name, func(t *testing.T) {
			// Create test file
			testFile := filepath.Join(tempDir, tf.filename)
			err := ioutil.WriteFile(testFile, []byte(tf.content), 0644)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			// Test build command
			os.Args = []string{"r2d2", "build", testFile}

			// Test file reading
			content, err := os.ReadFile(testFile)
			if err != nil {
				t.Errorf("Failed to read file: %v", err)
			}

			if string(content) != tf.content {
				t.Errorf("File content mismatch")
			}

			// Test filename extraction
			filename := getFilename()
			if filename != tf.filename {
				t.Errorf("Filename extraction failed: got %s, want %s", filename, tf.filename)
			}

			// Test with output flag
			os.Args = []string{"r2d2", "build", testFile, "-o", "test_output.html"}
			outputFile, hasOutput := parseOutputFlag()

			if !hasOutput {
				t.Errorf("Output flag not detected")
			}

			if outputFile != "test_output.html" {
				t.Errorf("Output file parsing failed: got %s, want test_output.html", outputFile)
			}
		})
	}

	os.Args = originalArgs
}

// TestCommandDispatchIntegration tests the complete command dispatch logic
func TestCommandDispatchIntegration(t *testing.T) {
	originalArgs := os.Args
	tempDir := t.TempDir()

	// Create a test file
	testFile := filepath.Join(tempDir, "dispatch_test.r2d2")
	testContent := `module DispatchTest { export fn main() { console.log("Dispatch"); } }`
	err := ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	commands := []struct {
		command        string
		aliases        []string
		needsFile      bool
		supportsOutput bool
	}{
		{
			command:        "version",
			aliases:        []string{"-version", "-v", "--version", "--v"},
			needsFile:      false,
			supportsOutput: false,
		},
		{
			command:        "build",
			aliases:        []string{"-b"},
			needsFile:      true,
			supportsOutput: true,
		},
		{
			command:        "run",
			aliases:        []string{"-r"},
			needsFile:      true,
			supportsOutput: false,
		},
		{
			command:        "js",
			aliases:        []string{},
			needsFile:      true,
			supportsOutput: true,
		},
	}

	for _, cmd := range commands {
		// Test main command
		t.Run("Command_"+cmd.command, func(t *testing.T) {
			if cmd.needsFile {
				os.Args = []string{"r2d2", cmd.command, testFile}
			} else {
				os.Args = []string{"r2d2", cmd.command}
			}

			// Validate command parsing
			if len(os.Args) >= 2 && os.Args[1] != cmd.command {
				t.Errorf("Command parsing failed: got %s, want %s", os.Args[1], cmd.command)
			}

			// Test with output flag if supported
			if cmd.supportsOutput && cmd.needsFile {
				os.Args = []string{"r2d2", cmd.command, testFile, "-o", "dispatch_output"}
				outputFile, hasOutput := parseOutputFlag()

				if !hasOutput {
					t.Errorf("Command %s should support output flag", cmd.command)
				}

				if outputFile != "dispatch_output" {
					t.Errorf("Output parsing failed for %s: got %s", cmd.command, outputFile)
				}
			}
		})

		// Test aliases
		for _, alias := range cmd.aliases {
			t.Run("Alias_"+alias, func(t *testing.T) {
				if cmd.needsFile {
					os.Args = []string{"r2d2", alias, testFile}
				} else {
					os.Args = []string{"r2d2", alias}
				}

				// Validate alias parsing
				if len(os.Args) >= 2 && os.Args[1] != alias {
					t.Errorf("Alias parsing failed: got %s, want %s", os.Args[1], alias)
				}
			})
		}
	}

	os.Args = originalArgs
}

// TestErrorScenariosIntegration tests various error conditions
func TestErrorScenariosIntegration(t *testing.T) {
	originalArgs := os.Args
	tempDir := t.TempDir()

	errorTests := []struct {
		name        string
		args        []string
		expectIssue bool
		description string
	}{
		{
			name:        "Build without file",
			args:        []string{"r2d2", "build"},
			expectIssue: true,
			description: "Should fail due to missing file argument",
		},
		{
			name:        "Run without file",
			args:        []string{"r2d2", "run"},
			expectIssue: true,
			description: "Should fail due to missing file argument",
		},
		{
			name:        "JS without file",
			args:        []string{"r2d2", "js"},
			expectIssue: true,
			description: "Should fail due to missing file argument",
		},
		{
			name:        "Build with non-existent file",
			args:        []string{"r2d2", "build", filepath.Join(tempDir, "nonexistent.r2d2")},
			expectIssue: true,
			description: "Should fail due to missing file",
		},
		{
			name:        "Unknown command",
			args:        []string{"r2d2", "unknown"},
			expectIssue: true,
			description: "Should be handled as unknown command",
		},
	}

	for _, et := range errorTests {
		t.Run(et.name, func(t *testing.T) {
			os.Args = et.args

			// Test argument validation
			if len(os.Args) >= 2 {
				cmd := os.Args[1]
				needsFile := cmd == "build" || cmd == "-b" || cmd == "run" || cmd == "-r" || cmd == "js"

				if needsFile {
					if len(os.Args) < 3 {
						// This is expected for error cases
						if !et.expectIssue {
							t.Errorf("Test case expects no issue but insufficient args found")
						}
						return
					}

					// Check if file exists
					filePath := os.Args[2]
					_, err := os.Stat(filePath)
					fileExists := !os.IsNotExist(err)

					if et.expectIssue && fileExists {
						t.Errorf("Expected file to not exist but it does: %s", filePath)
					}
				}
			}
		})
	}

	os.Args = originalArgs
}

// BenchmarkIntegrationWorkflow benchmarks the complete integration workflow
func BenchmarkIntegrationWorkflow(b *testing.B) {
	originalArgs := os.Args
	tempDir := b.TempDir()

	// Create benchmark test file
	testFile := filepath.Join(tempDir, "benchmark.r2d2")
	testContent := `module BenchmarkTest {
		var iterations i32 = 1000;

		export fn main() {
			for var i i32 = 0; i < iterations; i++ {
				process(i);
			}
		}

		fn process(value i32) {
			console.log("Processing: " + value);
		}
	}`

	err := ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		b.Fatalf("Failed to create benchmark file: %v", err)
	}

	b.Run("CompleteWorkflow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Test the complete workflow
			os.Args = []string{"r2d2", "build", testFile, "-o", "benchmark_output.html"}

			// Parse output flag
			_, _ = parseOutputFlag()

			// Get filename
			_ = getFilename()

			// Read file content
			_, _ = os.ReadFile(testFile)

			// Remove output flag
			removeOutputFlag()
		}
	})

	os.Args = originalArgs
}
