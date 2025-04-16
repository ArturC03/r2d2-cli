#!/usr/bin/env bash

# R2D2 Programming Language Installer
# Cross-platform installer for the R2D2 programming language

# ----------------------------------------
# Color definitions for better UI
# ----------------------------------------
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
MAGENTA='\033[0;35m'
CYAN='\033[0;36m'
BOLD='\033[1m'
RESET='\033[0m'

# ----------------------------------------
# Configuration variables
# ----------------------------------------
REPO_URL="https://github.com/ArturC03/r2d2-cli"
BINARY_NAME="r2d2"
TEMP_DIR="/tmp/r2d2-installer"
INSTALL_DIR="/usr/local/bin"
USER_INSTALL_DIR="$HOME/.local/bin"

# ----------------------------------------
# Helper functions
# ----------------------------------------
print_logo() {
  clear
  echo -e "${CYAN}"
  echo "  ███████████    ████████  ██████████    ████████ "
  echo " ░░███░░░░░███  ███░░░░███░░███░░░░███  ███░░░░███"
  echo "  ░███    ░███ ░░░    ░███ ░███   ░░███░░░    ░███"
  echo "  ░██████████     ███████  ░███    ░███   ███████ "
  echo "  ░███░░░░░███   ███░░░░   ░███    ░███  ███░░░░  "
  echo "  ░███    ░███  ███      █ ░███    ███  ███      █"
  echo "  █████   █████░██████████ ██████████  ░██████████"
  echo " ░░░░░   ░░░░░ ░░░░░░░░░░ ░░░░░░░░░░   ░░░░░░░░░░ "
  echo -e "${RESET}"
  echo -e "${BOLD}${MAGENTA} R2D2 Programming Language Installer ${RESET}"
  echo
}

spinner() {
  local pid=$1
  local delay=0.1
  local spinstr='|/-\'
  while ps -p $pid >/dev/null; do
    local temp=${spinstr#?}
    printf " [%c]  " "$spinstr"
    local spinstr=$temp${spinstr%"$temp"}
    sleep $delay
    printf "\b\b\b\b\b\b"
  done
  printf "    \b\b\b\b"
}

check_command() {
  command -v "$1" >/dev/null 2>&1
}

progress_bar() {
  local duration=$1
  local steps=20
  local delay=$(bc <<<"scale=2; $duration / $steps")

  echo -ne "${YELLOW}["
  for ((i = 0; i < steps; i++)); do
    sleep $delay
    echo -ne "="
  done
  echo -e "]${RESET}"
}

# ----------------------------------------
# OS Detection
# ----------------------------------------
detect_os() {
  if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Detect Linux distribution
    if [ -f /etc/os-release ]; then
      . /etc/os-release
      OS="$ID"
      OS_VER="$VERSION_ID"
    elif [ -f /etc/lsb-release ]; then
      . /etc/lsb-release
      OS="$DISTRIB_ID"
      OS_VER="$DISTRIB_RELEASE"
    elif [ -f /etc/debian_version ]; then
      OS="debian"
      OS_VER=$(cat /etc/debian_version)
    elif [ -f /etc/redhat-release ]; then
      OS="rhel"
    elif [ -f /etc/arch-release ]; then
      OS="arch"
    else
      OS="linux"
    fi
  elif [[ "$OSTYPE" == "darwin"* ]]; then
    OS="macos"
    OS_VER=$(sw_vers -productVersion)
  elif [[ "$OSTYPE" == "cygwin" ]] || [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
    OS="windows"
  else
    OS="unknown"
  fi
}

# ----------------------------------------
# Dependency Management
# ----------------------------------------
install_dependencies() {
  echo -e "${BLUE}${BOLD}[*] Checking and installing dependencies...${RESET}"

  case $OS in
  "ubuntu" | "debian" | "linuxmint" | "pop")
    echo -e "${YELLOW}[*] Detected Debian-based system: $OS $OS_VER${RESET}"
    echo -e "${YELLOW}[*] Installing dependencies with apt...${RESET}"
    sudo apt-get update
    sudo apt-get install -y git curl wget build-essential golang unzip
    curl -fsSL https://deno.land/install.sh | sh
    ;;
  "fedora" | "centos" | "rhel")
    echo -e "${YELLOW}[*] Detected RPM-based system: $OS $OS_VER${RESET}"
    echo -e "${YELLOW}[*] Installing dependencies with dnf/yum...${RESET}"
    if check_command dnf; then
      sudo dnf install -y git curl wget gcc gcc-c++ golang unzip
    else
      sudo yum install -y git curl wget gcc gcc-c++ golang unzip
    fi
    curl -fsSL https://deno.land/install.sh | sh
    ;;
  "arch" | "manjaro" | "endeavouros")
    echo -e "${YELLOW}[*] Detected Arch-based system: $OS${RESET}"
    echo -e "${YELLOW}[*] Installing dependencies with pacman...${RESET}"
    sudo pacman -Sy --noconfirm git curl wget base-devel go unzip
    sudo pacman -S --noconfirm deno
    ;;
  "opensuse" | "suse")
    echo -e "${YELLOW}[*] Detected openSUSE system${RESET}"
    echo -e "${YELLOW}[*] Installing dependencies with zypper...${RESET}"
    sudo zypper refresh
    sudo zypper install -y git curl wget gcc gcc-c++ go unzip
    curl -fsSL https://deno.land/install.sh | sh
    ;;
  "alpine")
    echo -e "${YELLOW}[*] Detected Alpine Linux${RESET}"
    echo -e "${YELLOW}[*] Installing dependencies with apk...${RESET}"
    sudo apk add git curl wget build-base go unzip
    curl -fsSL https://deno.land/install.sh | sh
    ;;
  "macos")
    echo -e "${YELLOW}[*] Detected macOS: $OS_VER${RESET}"
    if check_command brew; then
      echo -e "${YELLOW}[*] Installing dependencies with Homebrew...${RESET}"
      brew install git go deno
    else
      echo -e "${YELLOW}[*] Installing Homebrew first...${RESET}"
      /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
      brew install git go deno
    fi
    ;;
  "windows")
    echo -e "${YELLOW}[*] Detected Windows system${RESET}"
    echo -e "${YELLOW}[*] Please make sure you have Git, Go, and Deno installed${RESET}"
    echo -e "${YELLOW}[*] You can install them using scoop, chocolatey, or the official installers${RESET}"
    echo -e "${YELLOW}[*] Press Enter to continue if you have these dependencies installed${RESET}"
    read -p ""
    ;;
  *)
    echo -e "${RED}[!] Unsupported operating system: $OS${RESET}"
    echo -e "${YELLOW}[*] Please install Git, Go, and Deno manually${RESET}"
    echo -e "${YELLOW}[*] Press Enter to continue if you have these dependencies installed${RESET}"
    read -p ""
    ;;
  esac

  # Verify essential dependencies
  for cmd in git go; do
    if ! check_command $cmd; then
      echo -e "${RED}[!] Error: $cmd is not installed. Please install it manually and try again.${RESET}"
      exit 1
    fi
  done

  echo -e "${GREEN}[✓] Dependencies installed successfully${RESET}"
}

# ----------------------------------------
# Installation
# ----------------------------------------
install_r2d2() {
  echo -e "\n${BLUE}${BOLD}[*] Installing R2D2 Programming Language...${RESET}"

  # Create temporary directory
  echo -e "${YELLOW}[*] Creating temporary directory...${RESET}"
  mkdir -p "$TEMP_DIR"
  cd "$TEMP_DIR"

  # Clone repository
  echo -e "${YELLOW}[*] Cloning R2D2 repository...${RESET}"
  git clone "$REPO_URL" . &
  spinner $!

  if [ $? -ne 0 ]; then
    echo -e "${RED}[!] Error: Failed to clone the repository.${RESET}"
    exit 1
  fi

  # Build from source
  echo -e "${YELLOW}[*] Building R2D2 from source...${RESET}"
  echo -e "${YELLOW}[*] This may take a few minutes...${RESET}"
  go build . &
  spinner $!

  if [ $? -ne 0 ] || [ ! -f "$BINARY_NAME" ]; then
    echo -e "${RED}[!] Error: Failed to build R2D2.${RESET}"
    exit 1
  fi

  # Install binary
  echo -e "${YELLOW}[*] Installing R2D2 binary...${RESET}"

  # Check if we can install to system directory or need user directory
  if [ "$OS" == "windows" ]; then
    # For Windows/MSYS/Cygwin
    cp "$BINARY_NAME" "$BINARY_NAME.exe"
    mkdir -p "$HOME/bin"
    cp "$BINARY_NAME.exe" "$HOME/bin/"
    echo -e "${GREEN}[✓] R2D2 installed to $HOME/bin/${BINARY_NAME}.exe${RESET}"
    echo -e "${YELLOW}[*] Make sure $HOME/bin is in your PATH${RESET}"
  else
    # Try system-wide installation first, fall back to user directory
    if [ -w "$INSTALL_DIR" ] || [ "$(id -u)" -eq 0 ]; then
      cp "$BINARY_NAME" "$INSTALL_DIR/"
      chmod +x "$INSTALL_DIR/$BINARY_NAME"
      echo -e "${GREEN}[✓] R2D2 installed to $INSTALL_DIR/$BINARY_NAME${RESET}"
    else
      mkdir -p "$USER_INSTALL_DIR"
      cp "$BINARY_NAME" "$USER_INSTALL_DIR/"
      chmod +x "$USER_INSTALL_DIR/$BINARY_NAME"
      echo -e "${GREEN}[✓] R2D2 installed to $USER_INSTALL_DIR/$BINARY_NAME${RESET}"

      # Add to PATH if needed
      if [[ ":$PATH:" != *":$USER_INSTALL_DIR:"* ]]; then
        echo -e "${YELLOW}[*] Adding $USER_INSTALL_DIR to your PATH...${RESET}"

        # Detect shell and update appropriate config file
        SHELL_NAME=$(basename "$SHELL")
        if [ "$SHELL_NAME" == "bash" ]; then
          echo 'export PATH="$HOME/.local/bin:$PATH"' >>"$HOME/.bashrc"
          echo -e "${YELLOW}[*] Updated .bashrc, restart your terminal or run: source ~/.bashrc${RESET}"
        elif [ "$SHELL_NAME" == "zsh" ]; then
          echo 'export PATH="$HOME/.local/bin:$PATH"' >>"$HOME/.zshrc"
          echo -e "${YELLOW}[*] Updated .zshrc, restart your terminal or run: source ~/.zshrc${RESET}"
        elif [ "$SHELL_NAME" == "fish" ]; then
          echo 'set -gx PATH "$HOME/.local/bin" $PATH' >>"$HOME/.config/fish/config.fish"
          echo -e "${YELLOW}[*] Updated fish config, restart your terminal or run: source ~/.config/fish/config.fish${RESET}"
        else
          echo -e "${YELLOW}[*] Please add $USER_INSTALL_DIR to your PATH manually${RESET}"
        fi
      fi
    fi
  fi

  # Clean up
  echo -e "${YELLOW}[*] Cleaning up temporary files...${RESET}"
  cd
  rm -rf "$TEMP_DIR"
}

# ----------------------------------------
# Configuration
# ----------------------------------------
configure_r2d2() {
  echo -e "\n${BLUE}${BOLD}[*] Configuring R2D2...${RESET}"

  # Create config directory
  CONFIG_DIR="$HOME/.config/r2d2"
  mkdir -p "$CONFIG_DIR"

  # Create default config file if it doesn't exist
  if [ ! -f "$CONFIG_DIR/config.json" ]; then
    echo -e "${YELLOW}[*] Creating default configuration...${RESET}"
    cat >"$CONFIG_DIR/config.json" <<EOF
{
    "editor": "auto",
    "theme": "dark",
    "auto_update": true,
    "telemetry": false
}
EOF
    echo -e "${GREEN}[✓] Created default configuration at $CONFIG_DIR/config.json${RESET}"
  else
    echo -e "${GREEN}[✓] Configuration already exists at $CONFIG_DIR/config.json${RESET}"
  fi
}

# ----------------------------------------
# Post-installation
# ----------------------------------------
post_install() {
  echo -e "\n${BLUE}${BOLD}[*] Verifying installation...${RESET}"

  # Check if R2D2 is in PATH
  if check_command r2d2; then
    VERSION=$(r2d2 --version 2>/dev/null)
    echo -e "${GREEN}[✓] R2D2 installed successfully! Version: $VERSION${RESET}"
  else
    echo -e "${YELLOW}[!] R2D2 installed but may not be in your PATH.${RESET}"
    echo -e "${YELLOW}[!] You may need to restart your terminal or add R2D2 to your PATH manually.${RESET}"
  fi

  echo -e "\n${GREEN}${BOLD}=== R2D2 Programming Language Installation Complete ===${RESET}"
  echo -e "${CYAN}[*] Run 'r2d2 --help' to see available commands${RESET}"
  echo -e "${CYAN}[*] Documentation available at: https://github.com/ArturC03/r2d2-cli#readme${RESET}"
  echo -e "${CYAN}[*] Report issues at: https://github.com/ArturC03/r2d2-cli/issues${RESET}"
  echo -e "\n${GREEN}${BOLD}Happy coding with R2D2!${RESET}\n"
}

# ----------------------------------------
# Interactive Installation
# ----------------------------------------
interactive_install() {
  print_logo

  echo -e "${YELLOW}Welcome to the R2D2 Programming Language Installer!${RESET}"
  echo -e "${YELLOW}This script will install R2D2 and its dependencies on your system.${RESET}"
  echo

  # Detect OS
  echo -e "${BLUE}[*] Detecting operating system...${RESET}"
  detect_os
  echo -e "${GREEN}[✓] Detected: $OS ${OS_VER:-}${RESET}"
  echo

  # Installation options
  echo -e "${MAGENTA}${BOLD}Installation Options:${RESET}"
  echo -e "${CYAN}1) Express Install${RESET} - Install with default settings"
  echo -e "${CYAN}2) Custom Install${RESET} - Choose installation directory and components"
  echo -e "${CYAN}3) Exit${RESET} - Cancel installation"
  echo

  read -p "$(echo -e ${YELLOW}"Please select an option [1-3]: "${RESET})" choice

  case $choice in
  1)
    echo -e "${GREEN}[*] Starting express installation...${RESET}"
    install_dependencies
    echo -e "${YELLOW}[*] Installing R2D2...${RESET}"
    progress_bar 2
    install_r2d2
    configure_r2d2
    post_install
    ;;
  2)
    echo -e "${GREEN}[*] Starting custom installation...${RESET}"

    # Ask for dependency installation
    read -p "$(echo -e ${YELLOW}"Install dependencies? (y/n): "${RESET})" install_deps
    if [[ $install_deps =~ ^[Yy]$ ]]; then
      install_dependencies
    fi

    # Ask for installation directory
    read -p "$(echo -e ${YELLOW}"Installation directory [default: $INSTALL_DIR]: "${RESET})" custom_dir
    if [ -n "$custom_dir" ]; then
      INSTALL_DIR=$custom_dir
      USER_INSTALL_DIR=$custom_dir
    fi

    echo -e "${YELLOW}[*] Installing R2D2...${RESET}"
    progress_bar 2
    install_r2d2

    # Ask for configuration
    read -p "$(echo -e ${YELLOW}"Configure R2D2? (y/n): "${RESET})" configure
    if [[ $configure =~ ^[Yy]$ ]]; then
      configure_r2d2
    fi

    post_install
    ;;
  3)
    echo -e "${RED}Installation cancelled.${RESET}"
    exit 0
    ;;
  *)
    echo -e "${RED}Invalid option. Exiting.${RESET}"
    exit 1
    ;;
  esac
}

# ----------------------------------------
# Main Installation Script
# ----------------------------------------
main() {
  # Check if running with admin privileges on Windows
  if [ "$OS" == "windows" ] && ! net session &>/dev/null; then
    echo -e "${RED}[!] This script should be run with administrator privileges on Windows.${RESET}"
    echo -e "${RED}[!] Please run this script from an administrator command prompt.${RESET}"
    exit 1
  fi

  # Run interactive installer
  interactive_install
}

# Start the installer
detect_os
main
