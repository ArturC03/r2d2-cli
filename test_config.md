# R2D2-CLI Test Configuration

## Test Overview

This document describes the test suite for the r2d2-cli package. The tests are designed to cover all major functionality while excluding the parser module as requested.

## Test Files

### 1. `main_test.go`
Tests for core CLI functionality:
- **getFilename()**: File path extraction and validation
- **parseOutputFlag()**: Output flag parsing (-o flag)
- **removeOutputFlag()**: Output flag removal from arguments
- **Command line argument handling**: Various argument combinations
- **File path extraction**: Unix/Windows path handling
- **Command validation**: Valid/invalid command recognition
- **Benchmarks**: Performance testing for core functions

### 2. `commands_test.go`
Tests for command-related functionality:
- **ShowVersion()**: Version display functionality
- **UnknownCommand()**: Error handling for invalid commands
- **Message functions**: InfoMessage, ErrorMessage, HelpMessage
- **Version constant validation**: Semantic versioning checks
- **Style consistency**: Message formatting consistency
- **Special character handling**: Unicode, newlines, quotes
- **Benchmarks**: Performance testing for message functions

### 3. `styles_test.go`
Tests for styling and UI components:
- **Color constants**: Hex color validation
- **Adaptive colors**: Light/dark theme support
- **Tag styles**: Error, info, warning style rendering
- **UI component styles**: Container, heading, item styles
- **Style consistency**: Different style outputs
- **Empty content handling**: Edge case testing
- **Special character rendering**: Unicode, HTML, special chars
- **Color accessibility**: Basic contrast validation
- **Style composition**: Inheritance and base styles
- **Benchmarks**: Style rendering performance

### 4. `integration_test.go`
Integration tests for complete workflows:
- **Main function integration**: Command dispatch testing
- **File operations**: Reading, validation, error handling
- **Output flag integration**: Complete workflow testing
- **Command dispatch**: All command types and aliases
- **Error handling**: Various error scenarios
- **Concurrent access**: Race condition testing
- **Benchmarks**: End-to-end performance testing

## Test Coverage Areas

### ✅ Covered Functionality
- Command-line argument parsing
- File operations (reading, validation)
- Output flag handling (-o flag)
- Command validation and dispatch
- Error message formatting
- Style rendering and theming
- Version display
- Help system integration
- Edge cases and error scenarios
- Performance benchmarking
- Concurrent access patterns

### ❌ Excluded (As Requested)
- Parser module (`r2d2/parser/`)
- Actual R2D2 code compilation
- R2D2 code execution
- JavaScript transpilation (only interface testing)

## Running Tests

### Run All Tests
```bash
cd r2d2-cli
go test -v ./...
```

### Run Specific Test Files
```bash
# Main functionality
go test -v -run TestMain main_test.go

# Commands
go test -v -run TestCommand commands_test.go

# Styles
go test -v -run TestStyle styles_test.go

# Integration
go test -v -run TestIntegration integration_test.go
```

### Run Benchmarks
```bash
# All benchmarks
go test -v -bench=. ./...

# Specific benchmarks
go test -v -bench=BenchmarkGetFilename main_test.go
go test -v -bench=BenchmarkInfoMessage commands_test.go
go test -v -bench=BenchmarkErrorStyleRender styles_test.go
go test -v -bench=BenchmarkIntegration integration_test.go
```

### Test Coverage
```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# View coverage in terminal
go test -cover ./...
```

## Test Data Requirements

### Temporary Files
Tests automatically create temporary files in system temp directories using `t.TempDir()`. No manual cleanup required.

### Sample R2D2 Files
Tests use inline R2D2 code samples like:
```r2d2
module Test {
    export fn main() {
        console.log("Hello, R2D2!");
    }
}
```

### No External Dependencies
All tests are self-contained and don't require external R2D2 files or network access.

## Mock Strategy

Since we're excluding the parser and not testing actual R2D2 compilation:
- Tests validate function signatures and interfaces
- File content is tested for reading/validation only
- Command dispatch is tested without execution
- Error paths are tested with simulated conditions

## Performance Expectations

### Benchmark Targets
- `getFilename()`: < 1μs per operation
- `parseOutputFlag()`: < 5μs per operation
- Message formatting: < 10μs per operation
- Style rendering: < 50μs per operation
- File reading: < 1ms for typical files

### Memory Usage
- No memory leaks in repeated operations
- Efficient string handling
- Minimal allocations in hot paths

## Error Testing Strategy

### Covered Error Cases
- Insufficient command arguments
- Non-existent files
- Empty file paths
- Invalid command names
- Malformed output flags
- File permission issues
- Concurrent access scenarios

### Error Validation
- Error messages are informative
- Error codes are appropriate
- Recovery paths are tested
- Edge cases are handled gracefully

## CI/CD Integration

### Required Go Version
- Go 1.24.0 or later (as specified in go.mod)

### Test Commands for CI
```bash
# Install dependencies
go mod download

# Run tests with race detection
go test -race -v ./...

# Generate coverage
go test -coverprofile=coverage.out ./...

# Check coverage threshold (optional)
go tool cover -func=coverage.out | grep total
```

### Expected Test Results
- All tests should pass
- No race conditions detected
- Coverage should be > 80% for non-parser code
- Benchmarks should complete without timeouts

## Maintenance Notes

### Adding New Tests
1. Follow existing naming conventions
2. Use table-driven tests for multiple scenarios
3. Include both positive and negative test cases
4. Add benchmarks for performance-critical functions
5. Update this documentation

### Test Data Management
- Use `t.TempDir()` for temporary files
- Clean up resources in defer statements
- Avoid hardcoded paths or external dependencies

### Performance Monitoring
- Run benchmarks regularly
- Monitor for performance regressions
- Update benchmark targets as needed
