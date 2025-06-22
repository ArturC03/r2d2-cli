#!/bin/bash

# R2D2-CLI Test Runner Script
# This script provides convenient commands to run various test suites

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if Go is installed
check_go() {
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed or not in PATH"
        exit 1
    fi

    GO_VERSION=$(go version | cut -d' ' -f3)
    print_status "Using Go version: $GO_VERSION"
}

# Function to run all tests
run_all_tests() {
    print_status "Running all tests..."
    go test -v ./... 2>&1 | tee test_results.log

    if [ ${PIPESTATUS[0]} -eq 0 ]; then
        print_success "All tests passed!"
    else
        print_error "Some tests failed. Check test_results.log for details."
        exit 1
    fi
}

# Function to run tests with race detection
run_race_tests() {
    print_status "Running tests with race detection..."
    go test -race -v ./... 2>&1 | tee race_test_results.log

    if [ ${PIPESTATUS[0]} -eq 0 ]; then
        print_success "All race tests passed!"
    else
        print_error "Race conditions detected. Check race_test_results.log for details."
        exit 1
    fi
}

# Function to run specific test file
run_specific_test() {
    local test_file=$1
    if [ -z "$test_file" ]; then
        print_error "Please specify a test file"
        print_status "Available test files:"
        ls -1 *_test.go 2>/dev/null || print_warning "No test files found"
        exit 1
    fi

    if [ ! -f "$test_file" ]; then
        print_error "Test file '$test_file' not found"
        exit 1
    fi

    print_status "Running tests from $test_file..."
    go test -v -run ".*" "$test_file"
}

# Function to run benchmarks
run_benchmarks() {
    print_status "Running benchmark tests..."
    go test -v -bench=. -benchmem ./... 2>&1 | tee benchmark_results.log

    if [ ${PIPESTATUS[0]} -eq 0 ]; then
        print_success "Benchmarks completed successfully!"
        print_status "Results saved to benchmark_results.log"
    else
        print_error "Benchmark tests failed. Check benchmark_results.log for details."
        exit 1
    fi
}

# Function to generate coverage report
run_coverage() {
    print_status "Generating test coverage report..."

    # Run tests with coverage
    go test -coverprofile=coverage.out ./...

    if [ $? -eq 0 ]; then
        # Generate HTML coverage report
        go tool cover -html=coverage.out -o coverage.html

        # Show coverage summary
        echo ""
        print_status "Coverage Summary:"
        go tool cover -func=coverage.out | grep -E "(total|func)"

        print_success "Coverage report generated!"
        print_status "HTML report: coverage.html"
        print_status "Raw data: coverage.out"

        # Try to open HTML report if on macOS or Linux with display
        if command -v open &> /dev/null; then
            print_status "Opening coverage report in browser..."
            open coverage.html
        elif command -v xdg-open &> /dev/null; then
            print_status "Opening coverage report in browser..."
            xdg-open coverage.html
        else
            print_status "Open coverage.html in your browser to view the report"
        fi
    else
        print_error "Coverage generation failed"
        exit 1
    fi
}

# Function to run linting and formatting checks
run_lint() {
    print_status "Running linting and formatting checks..."

    # Check formatting
    UNFORMATTED=$(go fmt ./... | wc -l)
    if [ "$UNFORMATTED" -gt 0 ]; then
        print_warning "Some files were reformatted. Run 'go fmt ./...' to fix."
    else
        print_success "All files are properly formatted"
    fi

    # Run go vet
    print_status "Running go vet..."
    go vet ./...
    if [ $? -eq 0 ]; then
        print_success "No issues found by go vet"
    else
        print_error "go vet found issues"
        exit 1
    fi

    # Check for golint if available
    if command -v golint &> /dev/null; then
        print_status "Running golint..."
        golint ./...
    else
        print_warning "golint not found. Install with: go install golang.org/x/lint/golint@latest"
    fi
}

# Function to clean up test artifacts
cleanup() {
    print_status "Cleaning up test artifacts..."

    # Remove coverage files
    rm -f coverage.out coverage.html

    # Remove log files
    rm -f test_results.log race_test_results.log benchmark_results.log

    # Remove any temporary test files (if any exist)
    find . -name "*.tmp" -delete 2>/dev/null || true
    find . -name "test_*" -type f -delete 2>/dev/null || true

    print_success "Cleanup completed"
}

# Function to run quick smoke tests
run_quick() {
    print_status "Running quick smoke tests..."
    go test -short -v ./...

    if [ $? -eq 0 ]; then
        print_success "Quick tests passed!"
    else
        print_error "Quick tests failed"
        exit 1
    fi
}

# Function to show usage
show_usage() {
    echo "R2D2-CLI Test Runner"
    echo ""
    echo "Usage: $0 [command]"
    echo ""
    echo "Commands:"
    echo "  all         Run all tests (default)"
    echo "  race        Run tests with race detection"
    echo "  bench       Run benchmark tests"
    echo "  coverage    Generate test coverage report"
    echo "  lint        Run linting and formatting checks"
    echo "  quick       Run quick smoke tests"
    echo "  clean       Clean up test artifacts"
    echo "  specific    Run specific test file (requires filename)"
    echo "  help        Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0                      # Run all tests"
    echo "  $0 coverage            # Generate coverage report"
    echo "  $0 specific main_test.go  # Run main_test.go only"
    echo "  $0 bench               # Run benchmarks"
    echo ""
}

# Main script logic
main() {
    print_status "R2D2-CLI Test Runner Starting..."

    # Check Go installation
    check_go

    # Ensure we're in the right directory
    if [ ! -f "go.mod" ]; then
        print_error "go.mod not found. Please run this script from the r2d2-cli directory."
        exit 1
    fi

    # Download dependencies if needed
    print_status "Ensuring dependencies are available..."
    go mod download

    # Parse command line arguments
    case "${1:-all}" in
        "all")
            run_all_tests
            ;;
        "race")
            run_race_tests
            ;;
        "bench")
            run_benchmarks
            ;;
        "coverage")
            run_coverage
            ;;
        "lint")
            run_lint
            ;;
        "quick")
            run_quick
            ;;
        "clean")
            cleanup
            ;;
        "specific")
            run_specific_test "$2"
            ;;
        "help"|"-h"|"--help")
            show_usage
            ;;
        *)
            print_error "Unknown command: $1"
            echo ""
            show_usage
            exit 1
            ;;
    esac
}

# Run main function with all arguments
main "$@"
