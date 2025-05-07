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
REPO_URL="https://github.com/ArturC03/r2d2-cli.git"
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

  # Remove any existing clone to ensure the directory is empty
  rm -rf "$TEMP_DIR/*"

  # Clone repository
  echo -e "${YELLOW}[*] Cloning R2D2 repository...${RESET}"
  git clone "$REPO_URL" r2d2-cli &
  spinner $!
  cd r2d2-cli

  if [ $? -ne 0 ]; then
    echo -e "${RED}[!] Error: Failed to clone the repository.${RESET}"
    exit 1
  fi

  # Build from source
  echo -e "${YELLOW}[*] Building R2D2 from source...${RESET}"
  go build -o "$BINARY_NAME"

  if [ $? -ne 0 ]; then
    echo -e "${RED}[!] Error: Build failed.${RESET}"
    exit 1
  fi

  # Copy binary to appropriate location
  echo -e "${YELLOW}[*] Installing the binary...${RESET}"
  if [ -w "$INSTALL_DIR" ]; then
    sudo cp "$BINARY_NAME" "$INSTALL_DIR"
  else
    mkdir -p "$USER_INSTALL_DIR"
    cp "$BINARY_NAME" "$USER_INSTALL_DIR"
    echo -e "${YELLOW}[*] Installed to user local bin: $USER_INSTALL_DIR${RESET}"
  fi

  # Clean up
  rm -rf "$TEMP_DIR"

  echo -e "${GREEN}[✓] Installation successful!${RESET}"
}

# ----------------------------------------
# Main Execution
# ----------------------------------------
print_logo
detect_os
install_dependencies
install_r2d2
