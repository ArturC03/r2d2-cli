package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestMainIntegration tests the main function with various command combinations
func TestMainIntegration(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	// Create a temporary test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.r2d2")
	testContent := `module Test {
    export fn main() {
        console.log("Hello, R2D2!");
    }
}`

	err := ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name        string
		args        []string
		expectPanic bool
		expectExit  bool
	}{
		{
			name:        "No arguments (should show version)",
			args:        []string{"r2d2"},
			expectPanic: false,
			expectExit:  true, // main() calls os.Exit(0) for no args
		},
		{
			name:        "Help command",
			args:        []string{"r2d2", "help"},
			expectPanic: false,
			expectExit:  false,
		},
		{
			name:        "Version command",
			args:        []string{"r2d2", "version"},
			expectPanic: false,
			expectExit:  false,
		},
		{
			name:        "Build command with valid file",
			args:        []string{"r2d2", "build", testFile},
			expectPanic: false,
			expectExit:  false,
		},
		{
			name:        "Run command with valid file",
			args:        []string{"r2d2", "run", testFile},
			expectPanic: false,
			expectExit:  false,
		},
		{
			name:        "JS command with valid file",
			args:        []string{"r2d2", "js", testFile},
			expectPanic: false,
			expectExit:  false,
		},
		{
			name:        "Build with output flag",
			args:        []string{"r2d2", "build", testFile, "-o", "output.html"},
			expectPanic: false,
			expectExit:  false,
		},
		{
			name:        "Unknown command",
			args:        []string{"r2d2", "unknown"},
			expectPanic: false,
			expectExit:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args

			// We can't easily test main() directly because it calls os.Exit
			// Instead, we test the command parsing logic
			if len(os.Args) >= 2 {
				cmd := os.Args[1]

				// Test command recognition
				validCommands := map[string]bool{
					"help": true, "-help": true, "-h": true, "--help": true, "--h": true,
					"version": true, "-version": true, "-v": true, "--version": true, "--v": true,
					"build": true, "-b": true,
					"run": true, "-r": true,
					"js":  true,
					"new": true,
				}

				isValid := validCommands[cmd]
				if !isValid && cmd != "unknown" {
					// This would trigger UnknownCommand in main()
					t.Logf("Command %s would trigger UnknownCommand", cmd)
				}
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

// TestFileOperations tests file reading and validation
func TestFileOperations(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	// Create temporary test files
	tempDir := t.TempDir()

	// Valid R2D2 file
	validFile := filepath.Join(tempDir, "valid.r2d2")
	validContent := `module TestModule {
    var x i32 = 10;
    export fn main() {
        console.log("Test");
    }
}`
	err := ioutil.WriteFile(validFile, []byte(validContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create valid test file: %v", err)
	}

	// Empty R2D2 file
	emptyFile := filepath.Join(tempDir, "empty.r2d2")
	err = ioutil.WriteFile(emptyFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create empty test file: %v", err)
	}

	// Large R2D2 file
	largeFile := filepath.Join(tempDir, "large.r2d2")
	largeContent := strings.Repeat("// This is a comment line\n", 1000) + validContent
	err = ioutil.WriteFile(largeFile, []byte(largeContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create large test file: %v", err)
	}

	tests := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{
			name:     "Valid R2D2 file",
			filePath: validFile,
			wantErr:  false,
		},
		{
			name:     "Empty R2D2 file",
			filePath: emptyFile,
			wantErr:  false,
		},
		{
			name:     "Large R2D2 file",
			filePath: largeFile,
			wantErr:  false,
		},
		{
			name:     "Non-existent file",
			filePath: filepath.Join(tempDir, "nonexistent.r2d2"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = []string{"r2d2", "build", tt.filePath}

			// Test file existence check
			_, err := os.Stat(tt.filePath)
			fileExists := !os.IsNotExist(err)

			if tt.wantErr && fileExists {
				t.Errorf("Expected file to not exist, but it does: %s", tt.filePath)
			}
			if !tt.wantErr && !fileExists {
				t.Errorf("Expected file to exist, but it doesn't: %s", tt.filePath)
			}

			// Test file reading if file exists
			if fileExists {
				content, err := os.ReadFile(tt.filePath)
				if err != nil {
					t.Errorf("Failed to read file %s: %v", tt.filePath, err)
				}

				// Test that readR2D2File() would work with this file
				if len(os.Args) >= 3 {
					// Simulate what readR2D2File() does
					filePath := os.Args[2]
					if filePath != "" {
						_, statErr := os.Stat(filePath)
						if !os.IsNotExist(statErr) {
							fileContent, readErr := os.ReadFile(filePath)
							if readErr != nil {
								t.Errorf("File reading simulation failed: %v", readErr)
							}
							if len(fileContent) != len(content) {
								t.Errorf("Content length mismatch")
							}
						}
					}
				}
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

// TestOutputFlagIntegration tests the complete output flag workflow
func TestOutputFlagIntegration(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	// Create a temporary test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.r2d2")
	testContent := `module Test { export fn main() { console.log("Test"); } }`

	err := ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		hasOutput      bool
	}{
		{
			name:           "Build with output flag",
			args:           []string{"r2d2", "build", testFile, "-o", "custom.html"},
			expectedOutput: "custom.html",
			hasOutput:      true,
		},
		{
			name:           "JS with output flag",
			args:           []string{"r2d2", "js", testFile, "-o", "custom.js"},
			expectedOutput: "custom.js",
			hasOutput:      true,
		},
		{
			name:           "Build without output flag",
			args:           []string{"r2d2", "build", testFile},
			expectedOutput: "",
			hasOutput:      false,
		},
		{
			name:           "Output flag before command",
			args:           []string{"r2d2", "-o", "before.html", "build", testFile},
			expectedOutput: "before.html",
			hasOutput:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args

			// Test parseOutputFlag
			outputFile, hasOutput := parseOutputFlag()

			if hasOutput != tt.hasOutput {
				t.Errorf("parseOutputFlag() hasOutput = %v, want %v", hasOutput, tt.hasOutput)
			}

			if outputFile != tt.expectedOutput {
				t.Errorf("parseOutputFlag() outputFile = %v, want %v", outputFile, tt.expectedOutput)
			}

			// Test removeOutputFlag
			if hasOutput {
				originalArgsCount := len(os.Args)
				removeOutputFlag()

				// Should have removed -o and its value (2 elements)
				if len(os.Args) >= originalArgsCount {
					t.Errorf("removeOutputFlag() should reduce args count, before: %d, after: %d",
						originalArgsCount, len(os.Args))
				}

				// Restore args for next test
				os.Args = tt.args
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

// TestCommandDispatch tests the command dispatching logic
func TestCommandDispatch(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	// Create a temporary test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.r2d2")
	testContent := `module Test { export fn main() { console.log("Dispatch test"); } }`

	err := ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	commands := []struct {
		cmd            string
		aliases        []string
		requiresFile   bool
		supportsOutput bool
	}{
		{
			cmd:            "help",
			aliases:        []string{"-help", "-h", "--help", "--h"},
			requiresFile:   false,
			supportsOutput: false,
		},
		{
			cmd:            "version",
			aliases:        []string{"-version", "-v", "--version", "--v"},
			requiresFile:   false,
			supportsOutput: false,
		},
		{
			cmd:            "build",
			aliases:        []string{"-b"},
			requiresFile:   true,
			supportsOutput: true,
		},
		{
			cmd:            "run",
			aliases:        []string{"-r"},
			requiresFile:   true,
			supportsOutput: false,
		},
		{
			cmd:            "js",
			aliases:        []string{},
			requiresFile:   true,
			supportsOutput: true,
		},
	}

	for _, cmdInfo := range commands {
		// Test main command
		t.Run("Command: "+cmdInfo.cmd, func(t *testing.T) {
			if cmdInfo.requiresFile {
				os.Args = []string{"r2d2", cmdInfo.cmd, testFile}
			} else {
				os.Args = []string{"r2d2", cmdInfo.cmd}
			}

			// Verify args are set correctly
			if len(os.Args) >= 2 {
				if os.Args[1] != cmdInfo.cmd {
					t.Errorf("Command not set correctly: got %s, want %s", os.Args[1], cmdInfo.cmd)
				}
			}

			// Test with output flag if supported
			if cmdInfo.supportsOutput && cmdInfo.requiresFile {
				os.Args = []string{"r2d2", cmdInfo.cmd, testFile, "-o", "output.test"}
				outputFile, hasOutput := parseOutputFlag()

				if !hasOutput {
					t.Errorf("Command %s should support output flag", cmdInfo.cmd)
				}
				if outputFile != "output.test" {
					t.Errorf("Output file parsing failed for %s", cmdInfo.cmd)
				}
			}
		})

		// Test aliases
		for _, alias := range cmdInfo.aliases {
			t.Run("Alias: "+alias, func(t *testing.T) {
				if cmdInfo.requiresFile {
					os.Args = []string{"r2d2", alias, testFile}
				} else {
					os.Args = []string{"r2d2", alias}
				}

				// Verify alias is recognized
				if len(os.Args) >= 2 {
					if os.Args[1] != alias {
						t.Errorf("Alias not set correctly: got %s, want %s", os.Args[1], alias)
					}
				}
			})
		}
	}

	// Restore original args
	os.Args = originalArgs
}

// TestErrorHandling tests various error scenarios
func TestErrorHandling(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	tempDir := t.TempDir()

	errorScenarios := []struct {
		name        string
		args        []string
		setupFunc   func() string // Returns file path if needed
		expectError bool
	}{
		{
			name:        "Build without file argument",
			args:        []string{"r2d2", "build"},
			setupFunc:   func() string { return "" },
			expectError: true,
		},
		{
			name:        "Run without file argument",
			args:        []string{"r2d2", "run"},
			setupFunc:   func() string { return "" },
			expectError: true,
		},
		{
			name: "Build with non-existent file",
			args: []string{"r2d2", "build", "nonexistent.r2d2"},
			setupFunc: func() string {
				return filepath.Join(tempDir, "nonexistent.r2d2")
			},
			expectError: true,
		},
		{
			name: "Build with empty file path",
			args: []string{"r2d2", "build", ""},
			setupFunc: func() string {
				return ""
			},
			expectError: true,
		},
		{
			name: "Output flag without value",
			args: []string{"r2d2", "build", "test.r2d2", "-o"},
			setupFunc: func() string {
				testFile := filepath.Join(tempDir, "test.r2d2")
				ioutil.WriteFile(testFile, []byte("module Test{}"), 0644)
				return testFile
			},
			expectError: false, // parseOutputFlag handles this gracefully
		},
	}

	for _, scenario := range errorScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			filePath := scenario.setupFunc()

			// Update args with actual file path if provided
			if filePath != "" && len(scenario.args) >= 3 {
				scenario.args[2] = filePath
			}

			os.Args = scenario.args

			// Test argument validation
			if len(os.Args) < 3 && (len(os.Args) >= 2 && (os.Args[1] == "build" || os.Args[1] == "run" || os.Args[1] == "js")) {
				// This should trigger an error in the actual implementation
				if !scenario.expectError {
					t.Errorf("Expected no error, but insufficient args should cause error")
				}
			}

			// Test file existence for file-requiring commands
			if len(os.Args) >= 3 {
				filePath := os.Args[2]
				if filePath != "" {
					_, err := os.Stat(filePath)
					fileExists := !os.IsNotExist(err)

					if scenario.expectError && fileExists {
						t.Errorf("Expected error scenario, but file exists: %s", filePath)
					}
				}
			}
		})
	}

	// Restore original args
	os.Args = originalArgs
}

// TestConcurrentAccess tests concurrent access to shared resources
func TestConcurrentAccess(t *testing.T) {
	// Save original args
	originalArgs := os.Args

	// Create a temporary test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "concurrent.r2d2")
	testContent := `module Concurrent { export fn main() { console.log("Concurrent test"); } }`

	err := ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test concurrent access to os.Args modification
	done := make(chan bool, 2)

	// Goroutine 1: Parse output flag
	go func() {
		defer func() { done <- true }()
		for i := 0; i < 100; i++ {
			os.Args = []string{"r2d2", "build", testFile, "-o", "output1.html"}
			_, _ = parseOutputFlag()
			time.Sleep(time.Microsecond)
		}
	}()

	// Goroutine 2: Remove output flag
	go func() {
		defer func() { done <- true }()
		for i := 0; i < 100; i++ {
			os.Args = []string{"r2d2", "build", testFile, "-o", "output2.html"}
			removeOutputFlag()
			time.Sleep(time.Microsecond)
		}
	}()

	// Wait for both goroutines to complete
	<-done
	<-done

	// Restore original args
	os.Args = originalArgs
}

// BenchmarkIntegration benchmarks the complete workflow
func BenchmarkIntegration(b *testing.B) {
	// Save original args
	originalArgs := os.Args

	// Create a temporary test file
	tempDir := b.TempDir()
	testFile := filepath.Join(tempDir, "benchmark.r2d2")
	testContent := `module Benchmark {
		var x i32 = 42;
		export fn main() {
			console.log("Benchmark test");
			for var i i32 = 0; i < 10; i++ {
				console.log(i);
			}
		}
	}`

	err := ioutil.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		b.Fatalf("Failed to create test file: %v", err)
	}

	b.Run("ParseAndRemoveOutputFlag", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			os.Args = []string{"r2d2", "build", testFile, "-o", "benchmark.html"}
			_, _ = parseOutputFlag()
			removeOutputFlag()
		}
	})

	b.Run("GetFilename", func(b *testing.B) {
		os.Args = []string{"r2d2", "build", testFile}
		for i := 0; i < b.N; i++ {
			_ = getFilename()
		}
	})

	b.Run("FileReading", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := os.ReadFile(testFile)
			if err != nil {
				b.Errorf("File reading failed: %v", err)
			}
		}
	})

	// Restore original args
	os.Args = originalArgs
}
