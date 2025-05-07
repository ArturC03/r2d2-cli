#!/bin/bash

# R2D2 CLI Installer
# This script installs the R2D2 CLI tool on various operating systems

# Terminal colors and formatting
BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
WHITE='\033[1;37m'
BG_BLUE='\033[44m'
BG_GREEN='\033[42m'
BG_RED='\033[41m'
BG_YELLOW='\033[43m'
BOLD='\033[1m'
RESET='\033[0m'

# Function to display formatted messages
info() {
    echo -e "${WHITE}${BG_BLUE} INFO ${RESET} $1"
}

success() {
    echo -e "${WHITE}${BG_GREEN} SUCCESS ${RESET} $1"
}

error() {
    echo -e "${WHITE}${BG_RED} ERROR ${RESET} $1"
}

warning() {
    echo -e "${WHITE}${BG_YELLOW} WARNING ${RESET} $1"
}

# Detect OS
detect_os() {
    info "Detecting operating system..."
    
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        if [ -f /etc/debian_version ]; then
            OS="debian"
            PKG_MANAGER="apt-get"
        elif [ -f /etc/redhat-release ]; then
            OS="redhat"
            PKG_MANAGER="yum"
        elif [ -f /etc/arch-release ]; then
            OS="arch"
            PKG_MANAGER="pacman"
        elif [ -f /etc/SuSE-release ]; then
            OS="suse"
            PKG_MANAGER="zypper"
        else
            OS="linux"
            PKG_MANAGER="apt-get"
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        OS="macos"
        # Check if Homebrew is installed
        if command -v brew >/dev/null 2>&1; then
            PKG_MANAGER="brew"
        else
            error "Homebrew is not installed. Please install Homebrew first: https://brew.sh/"
            exit 1
        fi
    elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
        OS="windows"
        # Check if Chocolatey is installed
        if command -v choco >/dev/null 2>&1; then
            PKG_MANAGER="choco"
        else
            error "Chocolatey is not installed. Please install Chocolatey first: https://chocolatey.org/install"
            exit 1
        fi
    else
        error "Unsupported operating system: $OSTYPE"
        exit 1
    fi
    
    success "Detected OS: $OS with package manager: $PKG_MANAGER"
}

# Install dependencies based on OS
install_dependencies() {
    info "Installing dependencies (Git, Go, Deno)..."
    
    case $PKG_MANAGER in
        "apt-get")
            sudo apt-get update
            sudo apt-get install -y git golang
            # Install Deno
            curl -fsSL https://deno.land/install.sh | sh
            ;;
        "yum")
            sudo yum update -y
            sudo yum install -y git golang
            # Install Deno
            curl -fsSL https://deno.land/install.sh | sh
            ;;
        "pacman")
            sudo pacman -Syu --noconfirm
            sudo pacman -S --noconfirm git go
            # Install Deno
            curl -fsSL https://deno.land/install.sh | sh
            ;;
        "zypper")
            sudo zypper refresh
            sudo zypper install -y git go
            # Install Deno
            curl -fsSL https://deno.land/install.sh | sh
            ;;
        "brew")
            brew update
            brew install git go deno
            ;;
        "choco")
            choco install git golang deno -y
            ;;
        *)
            error "Unsupported package manager: $PKG_MANAGER"
            exit 1
            ;;
    esac
    
    # Verify installations
    if ! command -v git >/dev/null 2>&1; then
        error "Git installation failed."
        exit 1
    fi
    
    if ! command -v go >/dev/null 2>&1; then
        error "Go installation failed."
        exit 1
    fi
    
    # Add Deno to PATH if not already there
    if [[ "$OS" != "macos" && "$OS" != "windows" ]]; then
        export DENO_INSTALL="$HOME/.deno"
        export PATH="$DENO_INSTALL/bin:$PATH"
    fi
    
    if ! command -v deno >/dev/null 2>&1; then
        warning "Deno might not be in your PATH. You may need to manually add it."
        warning "Add these lines to your shell profile (.bashrc, .zshrc, etc.):"
        warning "export DENO_INSTALL=\"\$HOME/.deno\""
        warning "export PATH=\"\$DENO_INSTALL/bin:\$PATH\""
    else
        success "All dependencies installed successfully!"
    fi
}

# Clone and build the R2D2 CLI
clone_and_build() {
    info "Cloning the R2D2 CLI repository..."
    
    # Create a temporary directory for the build
    TMP_DIR=$(mktemp -d)
    cd "$TMP_DIR"
    
    # Clone the repository
    if ! git clone https://github.com/ArturC03/r2d2-cli.git; then
        error "Failed to clone the repository."
        exit 1
    fi
    
    cd r2d2-cli
    
    # Ensure go.mod exists, initialize if needed
    if [ ! -f "go.mod" ]; then
        info "Initializing Go module..."
        go mod init github.com/ArturC03/r2d2-cli
    fi
    
    # Ensure dependencies are properly listed in go.mod
    info "Tidying Go module dependencies..."
    go mod tidy
    
    info "Building the R2D2 CLI tool..."
    
    # Build the Go project with output name r2d2
    if ! go build -o r2d2 .; then
        error "Failed to build the R2D2 CLI tool."
        exit 1
    fi
    
    success "R2D2 CLI tool built successfully!"
    
    # Determine the bin directory to install to
    if [[ "$OS" == "windows" ]]; then
        BIN_DIR="$HOME/bin"
    else
        BIN_DIR="/usr/local/bin"
        # If user doesn't have permissions for /usr/local/bin, use $HOME/bin as fallback
        if [ ! -w "$BIN_DIR" ]; then
            BIN_DIR="$HOME/bin"
            mkdir -p "$BIN_DIR"
        fi
    fi
    
    info "Installing R2D2 CLI to $BIN_DIR..."
    
    # Move the binary to the bin directory
    if ! sudo mv r2d2 "$BIN_DIR/"; then
        error "Failed to move the binary to $BIN_DIR. Trying without sudo..."
        if ! mv r2d2 "$BIN_DIR/"; then
            error "Failed to install the R2D2 CLI tool."
            exit 1
        fi
    fi
    
    # Make sure the binary is executable
    if ! sudo chmod +x "$BIN_DIR/r2d2"; then
        if ! chmod +x "$BIN_DIR/r2d2"; then
            error "Failed to make the R2D2 CLI tool executable."
            exit 1
        fi
    fi
    
    # Clean up the temporary directory
    cd "$HOME"
    rm -rf "$TMP_DIR"
    
    success "R2D2 CLI tool installed successfully at $BIN_DIR/r2d2!"
    
    # Add the bin directory to PATH if it's not already there
    if [[ "$BIN_DIR" == "$HOME/bin" && "$PATH" != *"$BIN_DIR"* ]]; then
        warning "Please add $BIN_DIR to your PATH:"
        warning "export PATH=\"$BIN_DIR:\$PATH\""
        
        # Automatically add to shell profile
        SHELL_PROFILE=""
        if [ -f "$HOME/.bashrc" ]; then
            SHELL_PROFILE="$HOME/.bashrc"
        elif [ -f "$HOME/.zshrc" ]; then
            SHELL_PROFILE="$HOME/.zshrc"
        elif [ -f "$HOME/.profile" ]; then
            SHELL_PROFILE="$HOME/.profile"
        fi
        
        if [ -n "$SHELL_PROFILE" ]; then
            info "Attempting to add $BIN_DIR to PATH in $SHELL_PROFILE..."
            echo "export PATH=\"$BIN_DIR:\$PATH\"" >> "$SHELL_PROFILE"
            success "Added to $SHELL_PROFILE. Please restart your terminal or run: source $SHELL_PROFILE"
        fi
    fi
}

# Verify installation
verify_installation() {
    info "Verifying R2D2 CLI installation..."
    
    # Refresh PATH to include the binary location
    if [[ "$BIN_DIR" == "$HOME/bin" ]]; then
        export PATH="$BIN_DIR:$PATH"
    fi
    
    # Check if the r2d2 command is available
    if command -v r2d2 >/dev/null 2>&1; then
        success "R2D2 CLI tool installed and available in PATH!"
        info "You can now use the R2D2 CLI tool by typing 'r2d2' in your terminal."
    else
        warning "R2D2 CLI tool is installed but might not be in your PATH."
        warning "You can run it using the full path: $BIN_DIR/r2d2"
        warning "Or add $BIN_DIR to your PATH and restart your terminal."
    fi
}

# Main script execution
main() {
    clear

    # Print the ASCII art with bold green formatting
    echo -e ""
    echo -e "${BOLD}${GREEN}██████╗ ██████╗ ██████╗ ██████╗ ${RESET}"
    echo -e "${BOLD}${GREEN}██╔══██╗╚════██╗██╔══██╗╚════██╗${RESET}"
    echo -e "${BOLD}${GREEN}██████╔╝ █████╔╝██║  ██║ █████╔╝${RESET}"
    echo -e "${BOLD}${GREEN}██╔══██╗██╔═══╝ ██║  ██║██╔═══╝ ${RESET}"
    echo -e "${BOLD}${GREEN}██║  ██║███████╗██████╔╝███████╗${RESET}"
    echo -e "${BOLD}${GREEN}╚═╝  ╚═╝╚══════╝╚═════╝ ╚══════╝${RESET}"
    info "R2D2 CLI Installation Script"

    detect_os
    install_dependencies
    clone_and_build
    verify_installation
    
    echo ""
    echo -e "${BOLD}${GREEN}=======================================${RESET}"
    echo -e "${BOLD}${GREEN}   R2D2 CLI Installation Complete!    ${RESET}"
    echo -e "${BOLD}${GREEN}=======================================${RESET}"
}

# Execute the main function
main
